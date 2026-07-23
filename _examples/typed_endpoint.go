//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
)

func main() {
	log.SetFlags(0)

	es, err := elasticsearch.New()
	if err != nil {
		log.Fatalf("error creating the client: %s", err)
	}
	defer func() {
		closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := es.Close(closeCtx); err != nil {
			log.Fatalf("error closing the client: %s", err)
		}
	}()

	ctx := context.Background()

	infoRes, err := es.Info()
	if err != nil {
		log.Fatalf("functional Info: %s", err)
	}
	infoRes.Body.Close()

	typedSearch := search.New(es)

	res, err := typedSearch.
		Index("example").
		Query(esdsl.NewMatchQuery("title", "typed")).
		Do(ctx)
	if err != nil {
		log.Fatalf("typed Search: %s", err)
	}

	fmt.Printf("hits: %d\n", len(res.Hits.Hits))
	for _, hit := range res.Hits.Hits {
		fmt.Printf(" * id=%s source=%s\n", *hit.Id_, hit.Source_)
	}

	base, err := elasticsearch.NewBase()
	if err != nil {
		log.Fatalf("error creating the base client: %s", err)
	}
	defer func() {
		closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := base.Close(closeCtx); err != nil {
			log.Fatalf("error closing the base client: %s", err)
		}
	}()

	if _, err := search.New(base).
		Index("example").
		Query(esdsl.NewMatchQuery("title", "typed")).
		Do(ctx); err != nil {
		log.Fatalf("typed Search (NewBase): %s", err)
	}
}
