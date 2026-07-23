package commands

import (
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/elastic/go-elasticsearch/v9"

	"github.com/elastic/go-elasticsearch/v9/_examples/xkcdsearch"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search xkcd.com",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "\x1b[2m%s\x1b[0m\n", strings.Repeat("━", tWidth))
		fmt.Fprintf(os.Stdout, "\x1b[2m?q=\x1b[0m\x1b[1m%s\x1b[0m\n", strings.Join(args, " "))
		fmt.Fprintf(os.Stdout, "\x1b[2m%s\x1b[0m\n", strings.Repeat("━", tWidth))

		es, err := elasticsearch.NewTyped()
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;107;41mERROR: %s\x1b[0m\n", err)
		}
		defer func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := es.Close(ctx); err != nil {
				fmt.Fprintf(os.Stderr, "Error closing the client: %s\n", err)
			}
		}()

		config := xkcdsearch.StoreConfig{Client: es, IndexName: IndexName}
		store, err := xkcdsearch.NewStore(config)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;107;41mERROR: %s\x1b[0m\n", err)
			os.Exit(1)
		}
		search := Search{store: store, reHighlight: regexp.MustCompile("<em>(.+?)</em>")}

		results, err := search.getResults(strings.Join(args, " "))
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;107;41mERROR: %s\x1b[0m\n", err)
			os.Exit(1)
		}

		if results.Total < 1 {
			fmt.Fprintln(os.Stdout, "⨯ No results")
			fmt.Fprintf(os.Stdout, "\x1b[2m%s\x1b[0m\n", strings.Repeat("─", tWidth))
			os.Exit(0)
		}

		for _, result := range results.Hits {
			search.displayResult(os.Stdout, result)
		}
	},
}

type Search struct {
	store       *xkcdsearch.Store
	reHighlight *regexp.Regexp
}

func (s *Search) getResults(query string) (*xkcdsearch.SearchResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Search) displayResult(w io.Writer, hit *xkcdsearch.Hit) { _ = "STUB: not implemented"; return }

func (s *Search) highlightString(input string) string { _ = "STUB: not implemented"; return "" }
