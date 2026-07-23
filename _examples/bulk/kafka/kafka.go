//go:build ignore
// +build ignore

package main

import (
	"context"
	"expvar"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	_ "net/http/pprof"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esutil"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmelasticsearch"

	"github.com/elastic/go-elasticsearch/v9/_examples/bulk/kafka/consumer"
	"github.com/elastic/go-elasticsearch/v9/_examples/bulk/kafka/producer"
)

var (
	brokerURL string

	topicName  = "stocks"
	topicParts = 4
	msgRate    int

	indexName    = "stocks"
	numProducers = 1
	numConsumers = 4
	numIndexers  = 1
	flushBytes   = 0
	numWorkers   = 0
	indexerError error

	mapping = `{
  "mappings": {
    "properties": {
    	"time":     { "type": "date"    },
    	"symbol":   { "type": "keyword" },
    	"side":     { "type": "keyword" },
      "account":  { "type": "keyword" },
      "quantity": { "type": "long"    },
      "price":    { "type": "long"    },
      "amount":   { "type": "long"    }
      }
    }}}`
)

func init() {
	if v := os.Getenv("KAFKA_URL"); v != "" {
		brokerURL = v
	} else {
		brokerURL = "localhost:9092"
	}
	flag.IntVar(&msgRate, "rate", 1000, "Producer rate (msg/sec)")
	flag.IntVar(&numProducers, "producers", numProducers, "Number of producers")
	flag.IntVar(&numConsumers, "consumers", numConsumers, "Number of consumers")
	flag.IntVar(&numIndexers, "indexers", numIndexers, "Number of indexers")
	flag.Parse()
}

func main() {
	log.SetFlags(0)

	go func() { log.Println(http.ListenAndServe("localhost:6060", nil)) }()

	var (
		wg  sync.WaitGroup
		ctx = context.Background()

		producers []*producer.Producer
		consumers []*consumer.Consumer
		indexers  []esutil.BulkIndexer
	)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	go func() { <-done; log.Println(""); os.Exit(0) }()

	for i := 1; i <= numProducers; i++ {
		producers = append(producers,
			&producer.Producer{
				BrokerURL:   brokerURL,
				TopicName:   topicName,
				TopicParts:  topicParts,
				MessageRate: msgRate})
	}

	es, err := elasticsearch.New(
		elasticsearch.WithRetry(5, 502, 503, 504, 429),
		elasticsearch.WithTransportOptions(
			elastictransport.WithRetryBackoff(func(i int) time.Duration { return time.Duration(i) * 100 * time.Millisecond }),
			elastictransport.WithMetrics(),
			elastictransport.WithTransport(apmelasticsearch.WrapRoundTripper(http.DefaultTransport)),
		),
	)
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}

	expvar.Publish("go-elasticsearch", expvar.Func(func() interface{} { m, _ := es.Metrics(); return m }))

	res, err := es.Indices.Exists([]string{indexName})
	if err != nil {
		log.Fatalf("Error: Indices.Exists: %s", err)
	}
	res.Body.Close()
	if res.StatusCode == 404 {
		res, err := es.Indices.Create(
			indexName,
			es.Indices.Create.WithBody(strings.NewReader(mapping)),
			es.Indices.Create.WithWaitForActiveShards("1"),
		)
		if err != nil {
			log.Fatalf("Error: Indices.Create: %s", err)
		}
		if res.IsError() {
			log.Fatalf("Error: Indices.Create: %s", res)
		}
	}

	for i := 1; i <= numIndexers; i++ {
		idx, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Index:      indexName,
			Client:     es,
			NumWorkers: numWorkers,
			FlushBytes: int(flushBytes),

			OnFlushStart: func(ctx context.Context) context.Context {
				txn := apm.DefaultTracer.StartTransaction("Bulk", "indexing")
				return apm.ContextWithTransaction(ctx, txn)
			},
			OnFlushEnd: func(ctx context.Context) {
				apm.TransactionFromContext(ctx).End()
			},
			OnError: func(ctx context.Context, err error) {
				indexerError = err
				apm.CaptureError(ctx, err).Send()
			},
		})
		if err != nil {
			log.Fatalf("ERROR: NewBulkIndexer(): %s", err)
		}
		indexers = append(indexers, idx)
	}

	for i := 1; i <= numConsumers; i++ {
		consumers = append(consumers,
			&consumer.Consumer{
				BrokerURL: brokerURL,
				TopicName: topicName,
				Indexer:   indexers[i%numIndexers]})
	}

	reporter := time.NewTicker(500 * time.Millisecond)
	defer reporter.Stop()
	go func() {
		fmt.Printf("Initializing... producers=%d consumers=%d indexers=%d\n", numProducers, numConsumers, numIndexers)
		for {
			select {
			case <-reporter.C:
				fmt.Print(report(producers, consumers, indexers))
			}
		}
	}()
	errcleaner := time.NewTicker(10 * time.Second)
	defer errcleaner.Stop()
	go func() {
		for {
			select {
			case <-errcleaner.C:
				indexerError = nil
			}
		}
	}()

	if len(producers) > 0 {
		if err := producers[0].CreateTopic(ctx); err != nil {
			log.Fatalf("ERROR: Producer: %s", err)
		}
	}

	for _, c := range consumers {
		wg.Add(1)
		go func(c *consumer.Consumer) {
			defer wg.Done()
			if err := c.Run(ctx); err != nil {
				log.Fatalf("ERROR: Consumer: %s", err)
			}
		}(c)
	}

	time.Sleep(5 * time.Second)
	for _, p := range producers {
		wg.Add(1)
		go func(p *producer.Producer) {
			defer wg.Done()
			if err := p.Run(ctx); err != nil {
				log.Fatalf("ERROR: Producer: %s", err)
			}
		}(p)
	}

	wg.Wait()

	fmt.Print(report(producers, consumers, indexers))
}

func report(
	producers []*producer.Producer,
	consumers []*consumer.Consumer,
	indexers []esutil.BulkIndexer,
) string {
	_ = "STUB: not implemented"
	return ""
}
