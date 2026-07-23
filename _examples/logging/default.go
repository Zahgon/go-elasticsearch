//go:build logging_default
// +build logging_default

package main

import (
	"context"
	"log"
	"os"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
)

func main() {
	log.SetFlags(0)

	var es *elasticsearch.TypedClient

	es, _ = elasticsearch.NewTyped(elasticsearch.WithLogger(&elastictransport.TextLogger{Output: os.Stdout}))
	run(es, "Text")
	_ = es.Close(context.Background())

	es, _ = elasticsearch.NewTyped(elasticsearch.WithLogger(&elastictransport.ColorLogger{Output: os.Stdout}))
	run(es, "Color")
	_ = es.Close(context.Background())

	es, _ = elasticsearch.NewTyped(elasticsearch.WithLogger(&elastictransport.ColorLogger{
		Output:             os.Stdout,
		EnableRequestBody:  true,
		EnableResponseBody: true,
	}))
	run(es, "Request/Response Body")
	_ = es.Close(context.Background())

	es, _ = elasticsearch.NewTyped(elasticsearch.WithLogger(&elastictransport.CurlLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true}))
	run(es, "Curl")
	_ = es.Close(context.Background())

	es, _ = elasticsearch.NewTyped(elasticsearch.WithLogger(&elastictransport.JSONLogger{Output: os.Stdout}))
	run(es, "JSON")
	_ = es.Close(context.Background())
}

func run(es *elasticsearch.TypedClient, name string) { _ = "STUB: not implemented"; return }
