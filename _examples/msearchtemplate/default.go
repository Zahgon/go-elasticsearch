//go:build msearchtemplate_default
// +build msearchtemplate_default

package main

import (
	"context"
	"log"

	"github.com/elastic/go-elasticsearch/v9"
)

const demoIndex = "msearch-template-demo"

func main() {
	log.SetFlags(0)

	es, err := elasticsearch.New()
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}
	defer closeClient(es)

	typed, err := elasticsearch.NewTyped()
	if err != nil {
		log.Fatalf("Error creating typed client: %s", err)
	}
	defer closeTypedClient(typed)

	ctx := context.Background()
	setupDemoData(ctx, es)

	log.Println("=== MSearchTemplate using esapi (raw NDJSON) ===")
	runEsapiExample(ctx, es)

	log.Println()
	log.Println("=== MSearchTemplate using typedapi ===")
	runTypedAPIExample(ctx, typed)
}

func closeClient(es *elasticsearch.Client) { _ = "STUB: not implemented"; return }

func closeTypedClient(es *elasticsearch.TypedClient) { _ = "STUB: not implemented"; return }

func setupDemoData(ctx context.Context, es *elasticsearch.Client) {
	_ = "STUB: not implemented"
	return
}

func runEsapiExample(ctx context.Context, es *elasticsearch.Client) {
	_ = "STUB: not implemented"
	return
}

func runTypedAPIExample(ctx context.Context, typed *elasticsearch.TypedClient) {
	_ = "STUB: not implemented"
	return
}
