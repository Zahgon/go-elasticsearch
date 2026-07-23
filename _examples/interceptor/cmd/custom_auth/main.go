package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/_examples/interceptor/internal/fake"
)

func main() {

	srv := fake.NewServer(
		fake.WithLogFn(func(r *http.Request) {
			auth := r.Header.Get("Authorization")
			slog.Info("server received request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("authorization", auth),
			)
		}),
		fake.WithResponseBody([]byte(`{"cluster_name":"example"}`)),
		fake.WithHeaders(func(h http.Header) {
			h.Set("X-Elastic-Product", "Elasticsearch")
			h.Set("Content-Type", "application/json")
		}),
		fake.WithMiddleware(SPNEGOAuthMiddleware),
	)
	defer srv.Close()

	es, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses(srv.URL()),
		elasticsearch.WithTransportOptions(elastictransport.WithInterceptors(
			KerberosInterceptor(MockTokenProvider),
		)),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(">>> Sending request (interceptor will handle SPNEGO challenge)")
	resp, err := es.Info().Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf(">>> Response cluster_name: %s\n", resp.ClusterName)
}

func KerberosInterceptor(tokenProvider func() (string, error)) elastictransport.InterceptorFunc {
	_ = "STUB: not implemented"
	return *new(elastictransport.InterceptorFunc)
}

func SPNEGOAuthMiddleware(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func MockTokenProvider() (string, error) { _ = "STUB: not implemented"; return "", nil }
