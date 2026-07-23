//go:build logging_custom
// +build logging_custom

package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

type CustomLogger struct {
	zerolog.Logger
}

func (l *CustomLogger) LogRoundTrip(
	req *http.Request,
	res *http.Response,
	err error,
	start time.Time,
	dur time.Duration,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *CustomLogger) RequestBodyEnabled() bool { _ = "STUB: not implemented"; return false }

func (l *CustomLogger) ResponseBodyEnabled() bool { _ = "STUB: not implemented"; return false }

func main() {

	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Logger()

	es, _ := elasticsearch.NewTyped(elasticsearch.WithLogger(&CustomLogger{log}))
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := es.Close(ctx); err != nil {
			log.Fatal().Err(err).Msg("Error closing the client")
		}
	}()

	{
		ctx := context.Background()

		es.Delete("test", "1").Do(ctx)
		es.Exists("test", "1").Do(ctx)
		es.Index("test").
			Document(map[string]string{"title": "logging"}).
			Refresh(refresh.True).
			Do(ctx)

		es.Search().Q("{FAIL").Do(ctx)

		es.Search().
			Index("test").
			Query(esdsl.NewMatchQuery("title", "logging")).
			Size(1).
			Do(ctx)
	}
}
