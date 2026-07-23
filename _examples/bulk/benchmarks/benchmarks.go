//go:build ignore
// +build ignore

package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/valyala/fasthttp"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esutil"

	"github.com/elastic/go-elasticsearch/v9/_examples/bulk/benchmarks/runner"
)

type humanBytes uint64

func (b *humanBytes) String() string     { _ = "STUB: not implemented"; return "" }
func (b *humanBytes) Set(v string) error { _ = "STUB: not implemented"; return nil }

var (
	indexName     string
	datasetName   string
	numWorkers    int
	flushBytes    humanBytes
	numRuns       int
	numWarmupRuns int
	numItems      int
	numShards     int
	numReplicas   int
	wait          time.Duration
	useFasthttp   bool
	useEasyjson   bool
	mockserver    bool
	debug         bool
)

func init() {
	flag.StringVar(&indexName, "index", "test-bulk-benchmarks", "Index name")
	flag.StringVar(&datasetName, "dataset", "small", "Dataset to use for indexing")
	flag.IntVar(&numWorkers, "workers", runtime.NumCPU(), "Number of indexer workers")
	flag.Var(&flushBytes, "flush", "Flush threshold in bytes (default 3MB)")
	flag.IntVar(&numRuns, "runs", 10, "Number of runs")
	flag.IntVar(&numWarmupRuns, "warmup", 3, "Number of warmup runs")
	flag.IntVar(&numItems, "count", 100000, "Number of documents to generate")
	flag.IntVar(&numShards, "shards", 3, "Number of index shards")
	flag.IntVar(&numReplicas, "replicas", 0, "Number of index replicas (default 0)")
	flag.DurationVar(&wait, "wait", time.Second, "Wait duration between runs")
	flag.BoolVar(&useFasthttp, "fasthttp", false, "Use valyala/fasthttp for HTTP transport")
	flag.BoolVar(&useEasyjson, "easyjson", false, "Use mailru/easyjson for JSON decoding")
	flag.BoolVar(&mockserver, "mockserver", false, "Measure added, not flushed items")
	flag.BoolVar(&debug, "debug", false, "Enable logging output")
	flag.Parse()
}

func main() {
	log.SetFlags(0)

	indexName = indexName + "-" + datasetName
	if flushBytes < 1 {
		flushBytes = humanBytes(3e+6)
	}

	var opts []elasticsearch.Option

	if useFasthttp {
		opts = append(opts, elasticsearch.WithTransportOptions(elastictransport.WithTransport(&fasthttpTransport{})))
	}

	if debug {
		opts = append(opts, elasticsearch.WithLogger(&elastictransport.ColorLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true}))
	}

	es, _ := elasticsearch.New(opts...)

	runnerCfg := runner.Config{
		Client: es,

		IndexName:     indexName,
		DatasetName:   datasetName,
		NumShards:     numShards,
		NumReplicas:   numReplicas,
		NumItems:      numItems,
		NumRuns:       numRuns,
		NumWarmupRuns: numWarmupRuns,
		NumWorkers:    numWorkers,
		FlushBytes:    int(flushBytes),
		Wait:          wait,
		Mockserver:    mockserver,
	}

	if useEasyjson {
		runnerCfg.Decoder = easyjsonDecoder{}
	}

	runner, err := runner.NewRunner(runnerCfg)
	if err != nil {
		log.Fatalf("Error creating runner: %s", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() { <-done; log.Println("\r" + strings.Repeat("▁", 110)); runner.Report(); os.Exit(0) }()
	defer func() { log.Println(strings.Repeat("▁", 110)); runner.Report() }()

	log.Printf(
		"%s: run [%sx] warmup [%dx] shards [%d] replicas [%d] workers [%d] flush [%s] wait [%s]%s%s",
		datasetName,
		humanize.Comma(int64(numRuns)),
		numWarmupRuns,
		numShards,
		numReplicas,
		numWorkers,
		humanize.Bytes(uint64(flushBytes)),
		wait,
		func() string {
			if useFasthttp {
				return " fasthttp"
			}
			return ""
		}(),
		func() string {
			if useEasyjson {
				return " easyjson"
			}
			return ""
		}())
	log.Println(strings.Repeat("▔", 110))

	runner.Run()
}

type easyjsonDecoder struct{}

func (d easyjsonDecoder) UnmarshalFromReader(r io.Reader, blk *esutil.BulkIndexerResponse) error {
	_ = "STUB: not implemented"
	return nil
}

type fasthttpTransport struct{}

func (t *fasthttpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *fasthttpTransport) copyRequest(dst *fasthttp.Request, src *http.Request) *fasthttp.Request {
	_ = "STUB: not implemented"
	return nil
}

func (t *fasthttpTransport) copyResponse(dst *http.Response, src *fasthttp.Response) *http.Response {
	_ = "STUB: not implemented"
	return nil
}
