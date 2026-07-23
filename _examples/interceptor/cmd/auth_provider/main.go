package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"sync"

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

	authProvider := NewCredentialProvider("user1", "password1")

	es, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses(srv.URL()),
		elasticsearch.WithTransportOptions(elastictransport.WithInterceptors(
			DynamicAuthInterceptor(authProvider),
		)),
	)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	fmt.Println(">>> Sending request with initial credentials (user1)")
	_, _ = es.Info().Do(ctx)

	fmt.Println("\n>>> Rotating credentials to (user2)")
	authProvider.Update("user2", "password2")

	fmt.Println("\n>>> Sending request with rotated credentials (user2)")
	_, _ = es.Info().Do(ctx)
}

func DynamicAuthInterceptor(provider *CredentialProvider) elastictransport.InterceptorFunc {
	_ = "STUB: not implemented"
	return *new(elastictransport.InterceptorFunc)
}

type CredentialProvider struct {
	mu       sync.RWMutex
	username string
	password string
}

func NewCredentialProvider(username, password string) *CredentialProvider {
	_ = "STUB: not implemented"
	return nil
}

func (p *CredentialProvider) Update(username, password string) { _ = "STUB: not implemented"; return }

func (p *CredentialProvider) Get() (username, password string) {
	_ = "STUB: not implemented"
	return "", ""
}
