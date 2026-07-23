package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
)

func main() {
	log.SetFlags(0)

	es, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses("http://localhost:9200"),
		elasticsearch.WithTransportOptions(
			elastictransport.WithTransport(&http.Transport{
				MaxIdleConnsPerHost:   10,
				ResponseHeaderTimeout: time.Millisecond,
				DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
				TLSClientConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			}),
		),
	)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
	} else {
		log.Println(es.Info().Do(context.Background()))

	}
}
