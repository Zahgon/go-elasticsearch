//go:build ignore
// +build ignore

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
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
	defer infoRes.Body.Close()

	var info struct {
		Version struct {
			Number string `json:"number"`
		} `json:"version"`
	}
	if err := json.NewDecoder(infoRes.Body).Decode(&info); err != nil {
		log.Fatalf("decode: %s", err)
	}
	log.Printf("[functional] server version: %s", info.Version.Number)

	typed := elasticsearch.NewTypedFrom(es)

	typedInfo, err := typed.Info().Do(ctx)
	if err != nil {
		log.Fatalf("typed Info: %s", err)
	}
	log.Printf("[typed] server version: %s", typedInfo.Version.Int)

	log.Println(strings.Repeat("-", 37))

	if _, err := typed.Index("example").
		Id("1").
		Document(map[string]any{"title": "Hello, typed API"}).
		Do(ctx); err != nil {
		log.Fatalf("typed Index: %s", err)
	}

	searchRes, err := typed.Search().
		Index("example").
		Query(esdsl.NewMatchQuery("title", "typed")).
		Do(ctx)
	if err != nil {
		log.Fatalf("typed Search: %s", err)
	}
	for _, hit := range searchRes.Hits.Hits {
		fmt.Printf(" * id=%s source=%s\n", *hit.Id_, hit.Source_)
	}
}
