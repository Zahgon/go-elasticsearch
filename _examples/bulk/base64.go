//go:build bulk_base64
// +build bulk_base64

package main

import (
	"context"
	"flag"
	"log"
	"slices"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Doc struct {
	DocID string    `json:"docid"`
	Title string    `json:"title"`
	Text  string    `json:"text"`
	Emb   []float32 `json:"emb"`
}

type B64Doc struct {
	DocID string               `json:"docid"`
	Title string               `json:"title"`
	Text  string               `json:"text"`
	Emb   types.DenseVectorF32 `json:"emb"`
}

var (
	indexName string
	count     int
	batch     int
)

func init() {
	flag.StringVar(&indexName, "index", "openai-vector-bulk", "Index name")
	flag.IntVar(&count, "count", 20000, "Number of documents to index")
	flag.IntVar(&batch, "batch", 500, "Number of documents to send in one batch")
	flag.Parse()
}

func main() {
	log.SetFlags(0)

	log.Printf(
		"\x1b[1mBulk Base64\x1b[0m: documents [%s] batch size [%s]",
		humanize.Comma(int64(count)),
		humanize.Comma(int64(batch)),
	)
	log.Println("▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁")

	es, err := elasticsearch.NewTyped()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := es.Close(ctx); err != nil {
			log.Fatalf("Error closing the client: %s", err)
		}
	}()

	docs, err := loadB64Docs(count)
	if err != nil {
		log.Fatalf("Error loading documents: %s", err)
	}
	log.Printf("→ Loaded %s documents", humanize.Comma(int64(len(docs))))

	ctx := context.Background()
	if err := setupIndex(ctx, es, indexName); err != nil {
		log.Fatalf("Error setting up index: %s", err)
	}

	start := time.Now().UTC()
	for chunk := range slices.Chunk(docs, batch) {
		if _, err := ingestDocs(ctx, es, indexName, chunk); err != nil {
			log.Fatalf("Error executing bulk request: %s", err)
		}
	}

	log.Println("▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔")

	dur := time.Since(start)
	rate := float64(len(docs)) / dur.Seconds()
	log.Printf(
		"Successfully indexed [%s] documents in %s (%.0f docs/sec)",
		humanize.Comma(int64(len(docs))),
		dur.Truncate(time.Millisecond),
		rate,
	)
}

func loadB64Docs(count int) ([]B64Doc, error) { _ = "STUB: not implemented"; return nil, nil }

func setupIndex(ctx context.Context, es *elasticsearch.TypedClient, index string) error {
	_ = "STUB: not implemented"
	return nil
}

func ingestDocs(ctx context.Context, client *elasticsearch.TypedClient, index string, docs []B64Doc) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
