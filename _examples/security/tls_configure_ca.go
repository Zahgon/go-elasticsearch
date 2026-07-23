//go:build ca_from_tp
// +build ca_from_tp

package main

import (
	"context"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
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

	ca, err := ioutil.ReadFile(*cacert)
	if err != nil {
		log.Fatalf("ERROR: Unable to read CA from %q: %s", *cacert, err)
	}

	tp := http.DefaultTransport.(*http.Transport).Clone()

	if tp.TLSClientConfig.RootCAs, err = x509.SystemCertPool(); err != nil {
		log.Fatalf("ERROR: Problem adding system CA: %s", err)
	}

	if ok := tp.TLSClientConfig.RootCAs.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("ERROR: Problem adding CA from file %q", *cacert)
	}

	es, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses("https://localhost:9200"),
		elasticsearch.WithBasicAuth("elastic", *password),

		elasticsearch.WithTransportOptions(elastictransport.WithTransport(tp)),
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
