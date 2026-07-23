//go:build bulk_indexer
// +build bulk_indexer

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/dustin/go-humanize"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
	"github.com/elastic/go-elasticsearch/v9/esutil"
)

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Published time.Time `json:"published"`
	Author    Author    `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	indexName  string
	numWorkers int
	flushBytes int
	numItems   int
)

func init() {
	flag.StringVar(&indexName, "index", "test-bulk-example", "Index name")
	flag.IntVar(&numWorkers, "workers", runtime.NumCPU(), "Number of indexer workers")
	flag.IntVar(&flushBytes, "flush", 5e+6, "Flush threshold in bytes")
	flag.IntVar(&numItems, "count", 10000, "Number of documents to generate")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.SetFlags(0)

	var (
		articles        []*Article
		countSuccessful uint64

		res *esapi.Response
		err error
	)

	log.Printf(
		"\x1b[1mBulkIndexer\x1b[0m: documents [%s] workers [%d] flush [%s]",
		humanize.Comma(int64(numItems)), numWorkers, humanize.Bytes(uint64(flushBytes)))
	log.Println(strings.Repeat("▁", 65))

	retryBackoff := backoff.NewExponentialBackOff()

	es, err := elasticsearch.New(

		elasticsearch.WithRetry(5, 502, 503, 504, 429),

		elasticsearch.WithTransportOptions(
			elastictransport.WithRetryBackoff(func(i int) time.Duration {
				if i == 1 {
					retryBackoff.Reset()
				}
				return retryBackoff.NextBackOff()
			}),
		),
	)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         indexName,
		Client:        es,
		NumWorkers:    numWorkers,
		FlushBytes:    int(flushBytes),
		FlushInterval: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}

	names := []string{"Alice", "John", "Mary"}
	for i := 1; i <= numItems; i++ {
		articles = append(articles, &Article{
			ID:        i,
			Title:     strings.Join([]string{"Title", strconv.Itoa(i)}, " "),
			Body:      "Lorem ipsum dolor sit amet...",
			Published: time.Now().Round(time.Second).UTC().AddDate(0, 0, i),
			Author: Author{
				FirstName: names[rand.Intn(len(names))],
				LastName:  "Smith",
			},
		})
	}
	log.Printf("→ Generated %s articles", humanize.Comma(int64(len(articles))))

	if res, err = es.Indices.Delete([]string{indexName}, es.Indices.Delete.WithIgnoreUnavailable(true)); err != nil || res.IsError() {
		log.Fatalf("Cannot delete index: %s", err)
	}
	res.Body.Close()
	res, err = es.Indices.Create(indexName)
	if err != nil {
		log.Fatalf("Cannot create index: %s", err)
	}
	if res.IsError() {
		log.Fatalf("Cannot create index: %s", res)
	}
	res.Body.Close()

	start := time.Now().UTC()

	for _, a := range articles {

		data, err := json.Marshal(a)
		if err != nil {
			log.Fatalf("Cannot encode article %d: %s", a.ID, err)
		}

		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{

				Action: "index",

				DocumentID: strconv.Itoa(a.ID),

				Body: bytes.NewReader(data),

				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},

				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}

	}

	if err := bi.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	biStats := bi.Stats()

	log.Println(strings.Repeat("▔", 65))

	dur := time.Since(start)

	if biStats.NumFailed > 0 {
		log.Fatalf(
			"Indexed [%s] documents with [%s] errors in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			humanize.Comma(int64(biStats.NumFailed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	} else {
		log.Printf(
			"Successfully indexed [%s] documents in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	}
}
