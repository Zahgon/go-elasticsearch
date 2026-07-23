//go:build ca_from_client
// +build ca_from_client

package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
)

func main() {
	log.SetFlags(0)

	var (
		err error

		cacert   = flag.String("cacert", "certificates/ca/ca.crt", "Path to the file with certificate authority")
		password = flag.String("password", "elastic", "Elasticsearch password")
	)
	flag.Parse()

	cert, err := ioutil.ReadFile(*cacert)
	if err != nil {
		log.Fatalf("ERROR: Unable to read CA from %q: %s", *cacert, err)
	}

	es, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses("https://localhost:9200"),
		elasticsearch.WithBasicAuth("elastic", *password),

		elasticsearch.WithCACert(cert),
	)
	if err != nil {
		log.Fatalf("ERROR: Unable to create client: %s", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := es.Close(ctx); err != nil {
			log.Fatalf("Error closing the client: %s\n", err)
		}
	}()

	res, err := es.Info().Do(context.Background())
	if err != nil {
		log.Fatalf("ERROR: Unable to get response: %s", err)
	}

	log.Println(res)
}
