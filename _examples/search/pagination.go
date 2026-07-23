//go:build search_pagination
// +build search_pagination

package main

import (
	"context"
	"log"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
)

const (
	paginationIndex = "search-pagination-demo"
	totalDocs       = 25
	pageSize        = 10
)

type product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	log.SetFlags(0)

	es, err := elasticsearch.NewTyped()
	if err != nil {
		log.Fatalf("Error creating typed client: %s", err)
	}
	defer closePaginationClient(es)

	ctx := context.Background()
	setupPaginationData(ctx, es)

	log.Println("=== Pagination with from + size ===")
	runFromSizeExample(ctx, es)

	log.Println()
	log.Println("=== Pagination with PIT + search_after ===")
	runPITSearchAfterExample(ctx, es)
}

func closePaginationClient(es *elasticsearch.TypedClient) { _ = "STUB: not implemented"; return }

func setupPaginationData(ctx context.Context, es *elasticsearch.TypedClient) {
	_ = "STUB: not implemented"
	return
}

func runFromSizeExample(ctx context.Context, es *elasticsearch.TypedClient) {
	_ = "STUB: not implemented"
	return
}

func runPITSearchAfterExample(ctx context.Context, es *elasticsearch.TypedClient) {
	_ = "STUB: not implemented"
	return
}

func nextPITPage(req *search.Search, res *search.Response) *search.Search {
	_ = "STUB: not implemented"
	return nil
}
