package main

import (
	"context"
	"expvar"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"net/http"
	_ "net/http/pprof"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
)

var (
	tWidth, _, _ = terminal.GetSize(int(os.Stdout.Fd()))
)

func init() {
	runtime.SetMutexProfileFraction(10)
}

func main() {
	log.SetFlags(0)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	aborted := make(chan os.Signal)
	signal.Notify(aborted, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-aborted

		log.Println("\nDone!\n")
		os.Exit(0)
	}()

	log.Println(strings.Repeat("─", tWidth))
	log.Println("Open <http://localhost:6060/debug/vars> to see all exported variables.")
	log.Println(strings.Repeat("─", tWidth))

	go func() { log.Fatalln(http.ListenAndServe("localhost:6060", nil)) }()

	for i := 1; i <= 2; i++ {
		go func(i int) {
			log.Printf("==> Starting server on <localhost:1000%d>", i)
			if err := http.ListenAndServe(
				fmt.Sprintf("localhost:1000%d", i),
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "OK\n") }),
			); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Unable to start server: %s", err)
			}
		}(i)
	}

	es, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses(
			"http://localhost:10001",
			"http://localhost:10002",
			"http://localhost:10003",
		),
		elasticsearch.WithLogger(&elastictransport.ColorLogger{Output: os.Stdout}),
		elasticsearch.WithTransportOptions(
			elastictransport.WithDisableRetry(),
			elastictransport.WithDebugLogger(),

			elastictransport.WithMetrics(),
		),
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	expvar.Publish("go-elasticsearch", expvar.Func(func() interface{} { m, _ := es.Metrics(); return m }))

	ctx := context.Background()

	for {
		select {
		case t := <-ticker.C:

			go func() {
				_, _ = es.Info().Do(ctx)
			}()

			go func() {
				_, _ = es.Cluster.Health().Do(ctx)
			}()

			if t.Second()%5 == 0 {
				m, err := es.Metrics()
				if err != nil {
					log.Printf("\x1b[31;1mUnable to get metrics: %s\x1b[0m", err)
					continue
				}
				log.Println("███", fmt.Sprintf("\x1b[1m%s\x1b[0m", "Metrics"), strings.Repeat("█", tWidth-12))
				log.Printf(
					""+
						"    \x1b[2mRequests:   \x1b[0m %d\n"+
						"    \x1b[2mFailures:   \x1b[0m %d\n"+
						"    \x1b[2mConnections:\x1b[0m %s",
					m.Requests, m.Failures, m.Connections)
				log.Println(strings.Repeat("─", tWidth))
			}
		}
	}
}
