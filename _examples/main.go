package main

import (
	"context"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

type article struct {
	Title string `json:"title"`
}

func main() {
	log.SetFlags(0)

	es, err := elasticsearch.NewTyped()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := es.Close(ctx); err != nil {
			log.Fatalf("Error closing the client: %s\n", err)
		}
	}()

	ctx := context.Background()

	info, err := es.Info().Do(ctx)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", info.Version.Int)
	log.Println(strings.Repeat("~", 37))

	var wg sync.WaitGroup
	for i, title := range []string{"Test One", "Test Two"} {
		wg.Add(1)

		go func(i int, title string) {
			defer wg.Done()

			res, err := es.Index("test").
				Id(strconv.Itoa(i + 1)).
				Document(article{Title: title}).
				Refresh(refresh.True).
				Do(ctx)
			if err != nil {
				log.Printf("Error indexing document ID=%d: %s", i+1, err)
				return
			}

			log.Printf("%s; version=%d", res.Result, res.Version_)
		}(i, title)
	}
	wg.Wait()

	log.Println(strings.Repeat("-", 37))

	res, err := es.Search().
		Index("test").
		Query(esdsl.NewMatchQuery("title", "test")).
		TrackTotalHits(esdsl.NewTrackHits().Bool(true)).
		Pretty(true).
		Do(ctx)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	total := int64(0)
	if res.Hits.Total != nil {
		total = res.Hits.Total.Value
	}
	log.Printf("%d hits; took: %dms", total, res.Took)

	for _, hit := range res.Hits.Hits {
		id := ""
		if hit.Id_ != nil {
			id = *hit.Id_
		}
		log.Printf(" * ID=%s, %s", id, string(hit.Source_))
	}

	log.Println(strings.Repeat("=", 37))
}
