package main

import (
	"context"
	"log"
	"net/http/httptest"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {

	tp := newTracerProvider()
	defer func() { _ = tp.Shutdown(context.Background()) }()

	srv := newFakeES()
	defer srv.Close()

	captureSearchBody := true
	es, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses(srv.URL),
		elasticsearch.WithInstrumentation(
			elasticsearch.NewOpenTelemetryInstrumentation(tp, captureSearchBody),
		),
	)
	if err != nil {
		log.Fatalf("error creating client: %s", err)
	}

	ctx := context.Background()

	if _, err := es.Info().Do(ctx); err != nil {
		log.Fatalf("info: %s", err)
	}

	if _, err := es.Index("example").
		Id("1").
		Document(map[string]any{"title": "Hello, OpenTelemetry"}).
		Do(ctx); err != nil {
		log.Fatalf("index: %s", err)
	}

	if _, err := es.Search().
		Index("example").
		Query(esdsl.NewMatchQuery("title", "opentelemetry")).
		Do(ctx); err != nil {
		log.Fatalf("search: %s", err)
	}

	time.Sleep(100 * time.Millisecond)
}

func newTracerProvider() *sdktrace.TracerProvider { _ = "STUB: not implemented"; return nil }

func newFakeES() *httptest.Server { _ = "STUB: not implemented"; return nil }
