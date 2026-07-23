package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/_examples/interceptor/internal/fake"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

func main() {

	shutdown := initOTel()
	defer shutdown()

	srv := fake.NewServer(
		fake.WithStatusCode(http.StatusOK),
		fake.WithResponseBody([]byte(`{"cluster_name":"example","version":{"number":"9.2.0"}}`)),
		fake.WithHeaders(func(h http.Header) {
			h.Set("X-Elastic-Product", "Elasticsearch")
			h.Set("Content-Type", "application/json")
		}),
	)
	defer srv.Close()

	meter := otel.Meter("elasticsearch-client")
	requestCounter, _ := meter.Int64Counter("elasticsearch.client.requests",
		metric.WithDescription("Number of requests to Elasticsearch"),
		metric.WithUnit("{request}"),
	)
	requestDuration, _ := meter.Float64Histogram("elasticsearch.client.duration",
		metric.WithDescription("Duration of Elasticsearch requests"),
		metric.WithUnit("ms"),
	)

	tracer := otel.Tracer("elasticsearch-client")

	es, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses(srv.URL()),
		elasticsearch.WithTransportOptions(elastictransport.WithInterceptors(
			LoggingInterceptor(),
			MetricsInterceptor(requestCounter, requestDuration),
			TracingInterceptor(tracer),
		)),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(">>> Sending requests to demonstrate observability interceptors")
	fmt.Println()

	for i := 1; i <= 3; i++ {
		fmt.Printf("--- Request %d ---\n", i)
		_, _ = es.Info().Do(context.Background())
		fmt.Println()
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println(">>> Waiting for metrics to flush...")
	time.Sleep(2 * time.Second)
}

func initOTel() func() { _ = "STUB: not implemented"; return nil }

func LoggingInterceptor() elastictransport.InterceptorFunc {
	_ = "STUB: not implemented"
	return *new(elastictransport.InterceptorFunc)
}

func MetricsInterceptor(counter metric.Int64Counter, histogram metric.Float64Histogram) elastictransport.InterceptorFunc {
	_ = "STUB: not implemented"
	return *new(elastictransport.InterceptorFunc)
}

func TracingInterceptor(tracer trace.Tracer) elastictransport.InterceptorFunc {
	_ = "STUB: not implemented"
	return *new(elastictransport.InterceptorFunc)
}
