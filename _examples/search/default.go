//go:build search_default
// +build search_default

package main

import (
	"context"
	"log"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
)

const demoIndex = "search-demo"

type article struct {
	Title    string `json:"title"`
	Category string `json:"category"`
}

func main() {
	log.SetFlags(0)

	typed, err := elasticsearch.NewTyped()
	if err != nil {
		log.Fatalf("Error creating typed client: %s", err)
	}
	defer closeClient(typed)

	ctx := context.Background()
	setupDemoData(ctx, typed)

	log.Println("=== Search using esdsl builders ===")
	runEsdslExample(ctx, typed)

	log.Println()
	log.Println("=== Search using raw types.* structs ===")
	runTypesExample(ctx, typed)
}

func closeClient(es *elasticsearch.TypedClient) { _ = "STUB: not implemented"; return }

func setupDemoData(ctx context.Context, es *elasticsearch.TypedClient) {
	_ = "STUB: not implemented"
	return
}

func runEsdslExample(ctx context.Context, es *elasticsearch.TypedClient) {
	_ = "STUB: not implemented"
	return
}

func runTypesExample(ctx context.Context, es *elasticsearch.TypedClient) {
	_ = "STUB: not implemented"
	return
}

func printHits(res *search.Response) { _ = "STUB: not implemented"; return }
