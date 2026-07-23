package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/_examples/interceptor/internal/fake"
	"github.com/elastic/go-elasticsearch/v9/_examples/interceptor/internal/redact"
)

func main() {

	srv := fake.NewServer(
		fake.WithLogFn(func(r *http.Request) {
			username, password, _ := redact.BasicAuth(r)
			slog.Info("server received request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("username", username),
				slog.String("password", password),
			)
		}),
		fake.WithStatusCode(http.StatusOK),
		fake.WithResponseBody([]byte(`{"cluster_name":"example"}`)),
		fake.WithHeaders(func(h http.Header) {
			h.Set("X-Elastic-Product", "Elasticsearch")
			h.Set("Content-Type", "application/json")
		}),
	)
	defer srv.Close()

	es, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses(srv.URL()),
		elasticsearch.WithBasicAuth("default_user", "default_password"),
		elasticsearch.WithTransportOptions(elastictransport.WithInterceptors(
			ContextAuthInterceptor(),
		)),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(">>> Sending request with default credentials")
	_, _ = es.Info().Do(context.Background())

	fmt.Println("\n>>> Sending request with context override (tenant_a)")
	ctx := WithBasicAuth(context.Background(), "tenant_a", "tenant_a_secret")
	_, _ = es.Info().Do(ctx)

	fmt.Println("\n>>> Sending request with context override (tenant_b)")
	ctx = WithBasicAuth(context.Background(), "tenant_b", "tenant_b_secret")
	_, _ = es.Info().Do(ctx)

	fmt.Println("\n>>> Sending request with default credentials again")
	_, _ = es.Info().Do(context.Background())
}

type basicAuthKey struct{}

type basicAuthValue struct {
	username string
	password string
}

func WithBasicAuth(ctx context.Context, username, password string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func ContextAuthInterceptor() elastictransport.InterceptorFunc {
	_ = "STUB: not implemented"
	return *new(elastictransport.InterceptorFunc)
}
