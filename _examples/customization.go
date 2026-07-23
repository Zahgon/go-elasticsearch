//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
)

func CountingInterceptor(count *atomic.Uint64) elastictransport.InterceptorFunc {
	_ = "STUB: not implemented"
	return *new(elastictransport.InterceptorFunc)
}

func main() {
	var wg sync.WaitGroup
	var count atomic.Uint64

	es, _ := elasticsearch.NewTyped(
		elasticsearch.WithTransportOptions(elastictransport.WithInterceptors(
			CountingInterceptor(&count),
		)),
	)

	ctx := context.Background()
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			es.Info().Do(ctx)
		}()
	}
	wg.Wait()

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%80s\n", fmt.Sprintf("Total Requests: %d", count.Load()))
}
