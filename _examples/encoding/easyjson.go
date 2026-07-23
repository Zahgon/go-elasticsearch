package main

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
	"github.com/fatih/color"
	"github.com/mailru/easyjson"

	"github.com/elastic/go-elasticsearch/v9/_examples/encoding/model"
)

var (
	out       = color.New(color.Reset)
	faint     = color.New(color.Faint)
	bold      = color.New(color.Bold)
	red       = color.New(color.FgRed)
	boldGreen = color.New(color.Bold, color.FgGreen)
	boldRed   = color.New(color.Bold, color.FgRed)

	articles []model.Article
	fnames   []string
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	es, err := elasticsearch.New()
	if err != nil {
		fmt.Printf("Error creating the client: %s\n", err)
		os.Exit(2)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := es.Close(ctx); err != nil {
			fmt.Printf("Error closing the client: %s\n", err)
		}
	}()

	fnames = []string{"Alice", "John", "Mary"}

	for i, title := range []string{"One", "Two", "Three", "Four", "Five"} {
		articles = append(articles,
			model.Article{
				ID:        uint(i + 1),
				Title:     "Test " + title,
				Body:      "Lorem ipsum dolor sit amet, consectetur adipisicing elit",
				Published: time.Now().AddDate(i, 0, 0),
				Author: &model.Author{
					FirstName: fnames[rand.Intn(len(fnames))],
					LastName:  "Smith",
				},
			})
	}

	faint.Println("Indexing articles...")
	faint.Println(strings.Repeat("━", 80))

	var b bytes.Buffer
	for _, a := range articles {
		b.Reset()

		if _, err := easyjson.MarshalToWriter(a, &b); err != nil {
			red.Println("Error decoding response", err)
			continue
		}

		res, err := es.Index(
			"articles",
			bytes.NewReader(b.Bytes()),
			es.Index.WithDocumentID(strconv.Itoa(int(a.ID))),
		)
		if err != nil {
			red.Printf("Error indexing article: %s\n", err)
			continue
		}
		defer res.Body.Close()

		if res.IsError() {
			printErrorResponse(res)
			continue
		}

		var ir model.IndexResponse
		if err := easyjson.UnmarshalFromReader(res.Body, &ir); err != nil {
			red.Println("Error decoding response", err)
			continue
		}

		boldGreen.Printf("[%s] ", res.Status())
		fmt.Println(
			faint.Sprint("result=")+out.Sprint(ir.Result),
			faint.Sprint("index=")+out.Sprint(ir.Index),
			faint.Sprint("ID=")+out.Sprint(ir.ID),
			faint.Sprint("version=")+out.Sprint(ir.Version),
		)
	}

	es.Indices.Refresh(es.Indices.Refresh.WithIndex("articles"))

	faint.Println("\nSearching articles...")
	faint.Println(strings.Repeat("━", 80))

	res, err := es.Search(
		es.Search.WithIndex("articles"),
		es.Search.WithQuery("one OR two"),
	)
	if err != nil {
		red.Printf("Error searching articles: %s\n", err)
		os.Exit(2)
	}
	defer res.Body.Close()

	if res.IsError() {
		printErrorResponse(res)
		os.Exit(2)
	}

	var sr model.SearchResponse
	if err := easyjson.UnmarshalFromReader(res.Body, &sr); err != nil {
		red.Println("Error decoding response", err)
		os.Exit(2)
	}

	faint.Printf("[%s] took=%d total=%d\n", res.Status(), sr.Took, sr.Hits.Total.Value)
	faint.Println(strings.Repeat("─", 80))

	for _, h := range sr.Hits.Hits {
		fmt.Println(
			out.Sprintf("%s,", strings.Join([]string{h.Source.Author.FirstName, h.Source.Author.LastName}, " ")),
			bold.Sprintf("%s", h.Source.Title),
			out.Sprintf("(%d)", h.Source.Published.Year()),
		)
		faint.Println(strings.Repeat("─", 80))
	}
}

func printErrorResponse(res *esapi.Response) { _ = "STUB: not implemented"; return }
