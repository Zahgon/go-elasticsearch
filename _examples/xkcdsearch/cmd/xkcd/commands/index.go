package commands

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/elastic/go-elasticsearch/v9"

	"github.com/elastic/go-elasticsearch/v9/_examples/xkcdsearch"
)

var (
	indexSetup     bool
	crawlerWorkers int
)

func init() {
	rootCmd.AddCommand(indexCmd)
	indexCmd.Flags().BoolVar(&indexSetup, "setup", false, "Create Elasticsearch index")
	indexCmd.Flags().IntVar(&crawlerWorkers, "workers", 25, "Number of concurrent workers")
}

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Index xkcd.com into Elasticsearch",
	Run: func(cmd *cobra.Command, args []string) {
		crawler := Crawler{
			log: zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).
				Level(func() zerolog.Level {
					if os.Getenv("DEBUG") != "" {
						return zerolog.DebugLevel
					} else {
						return zerolog.InfoLevel
					}
				}()).
				With().
				Timestamp().
				Logger(),

			workers: crawlerWorkers,
			queue:   make(chan string, crawlerWorkers),
			reURL:   regexp.MustCompile(`https\://xkcd\.com/(?P<ID>\d+)/.*\.json`),
			nextURL: func(c *Crawler, u string) string {
				id, err := c.documentIDFromURL(u)
				if err != nil || id <= 1 {
					return ""
				}
				return fmt.Sprintf("https://xkcd.com/%d/info.0.json", id-1)
			},

			StartURL: "https://xkcd.com/info.0.json",
		}

		es, err := elasticsearch.NewTyped()
		if err != nil {
			crawler.log.Fatal().Err(err).Msg("Error creating Elasticsearch client")
		}
		defer func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := es.Close(ctx); err != nil {
				crawler.log.Fatal().Err(err).Msg("Error closing the client")
			}
		}()

		config := xkcdsearch.StoreConfig{Client: es, IndexName: IndexName}
		store, err := xkcdsearch.NewStore(config)
		if err != nil {
			crawler.log.Fatal().Err(err).Msg("Cannot create store")
		}
		crawler.store = store

		if indexSetup {
			crawler.log.Info().Msg("Creating index with mapping")
			if err := crawler.setupIndex(); err != nil {
				crawler.log.Fatal().Err(err).Msg("Cannot create Elasticsearch index")
			}
		}

		crawler.log.Info().Msgf("Starting the crawl with %d workers at <%s>", crawler.workers, crawler.StartURL)
		crawler.Run()
	},
}

type Crawler struct {
	store *xkcdsearch.Store
	log   zerolog.Logger

	workers int
	queue   chan string
	reURL   *regexp.Regexp
	nextURL func(crawler *Crawler, currentURL string) (nextURL string)

	StartURL string
}

func (c *Crawler) Run() { _ = "STUB: not implemented"; return }

func (c *Crawler) ProcessURL(url string) (doc xkcdsearch.Document) {
	_ = "STUB: not implemented"
	return *new(xkcdsearch.Document)
}

func (c *Crawler) NextURL(url string) string { _ = "STUB: not implemented"; return "" }

func (c *Crawler) documentIDFromURL(url string) (id int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Crawler) loadURL(url string) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Crawler) existsDocument(id string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Crawler) storeDocument(doc *xkcdsearch.Document) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Crawler) processResponse(res *http.Response) (xkcdsearch.Document, error) {
	_ = "STUB: not implemented"
	return *new(xkcdsearch.Document), nil
}

func (c *Crawler) setupIndex() error { _ = "STUB: not implemented"; return nil }
