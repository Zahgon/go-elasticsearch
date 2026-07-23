//go:build ignore
// +build ignore

package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
)

const port = "9209"

type ExtendedClient struct {
	*elasticsearch.Client
	Custom *ExtendedAPI
}

type ExtendedAPI struct {
	*elasticsearch.Client
}

func (c *ExtendedAPI) Example() (*esapi.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	log.SetFlags(0)

	started := make(chan bool)

	go startServer(started)

	esclient, err := elasticsearch.New(
		elasticsearch.WithAddresses("http://localhost:"+port),
		elasticsearch.WithLogger(&elastictransport.ColorLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true}),
	)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := esclient.Close(ctx); err != nil {
			log.Fatalf("Error closing the client: %s\n", err)
		}
	}()

	es := ExtendedClient{Client: esclient, Custom: &ExtendedAPI{esclient}}
	<-started

	es.Cat.Health()

	res, err := es.Custom.Example()
	if err != nil {
		log.Fatalf("Error calling custom API: %s", err)
	}
	defer res.Body.Close()
	log.Println(res.Status())
}

func startServer(started chan<- bool) { _ = "STUB: not implemented"; return }
