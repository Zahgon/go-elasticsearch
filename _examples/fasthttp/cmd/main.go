package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/_examples/fasthttp"
)

var (
	count     int
	durations []time.Duration
)

func init() {
	if envCount := os.Getenv("COUNT"); envCount != "" {
		i, err := strconv.Atoi(envCount)
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}
		count = i
	} else {
		count = 1000
	}
}

func main() {
	log.SetFlags(0)

	es, err := elasticsearch.New(
		elasticsearch.WithTransportOptions(elastictransport.WithTransport(&fasthttp.Transport{})),
	)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := es.Close(ctx); err != nil {
			fmt.Printf("Error closing the client: %s\n", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	go func() {
		res, err := es.Search(
			es.Search.WithBody(strings.NewReader(`{"query":{"match":{"title":"foo"}}}`)),
			es.Search.WithPretty(),
		)
		cancel()
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()
		if res.IsError() {
			log.Fatalf("Error response: %s", res)
		}
	}()

	select {
	case <-ctx.Done():
		if ctx.Err() != context.Canceled {
			log.Fatalf("Timeout: %s", ctx.Err())
		}
	}

	t := time.Now()
	for i := 0; i < count; i++ {
		t0 := time.Now()
		res, err := es.Info()
		durations = append(durations, time.Now().Sub(t0))
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		res.Body.Close()
	}

	sorted := append(make([]time.Duration, 0), durations...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	log.Printf(
		"%d requests in %.2f (%.1f req/s) | min: %s / max: %s / mean: %s",
		count,
		time.Now().Sub(t).Seconds(),
		float64(count)/time.Now().Sub(t).Seconds(),
		sorted[0],
		sorted[(len(sorted)-1)],
		sorted[(len(sorted)-1)/2],
	)
}
