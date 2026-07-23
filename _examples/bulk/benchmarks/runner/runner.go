package runner

import (
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esutil"
)

func NewRunner(cfg Config) (*Runner, error) { _ = "STUB: not implemented"; return nil, nil }

type Runner struct {
	config Config

	doc           []byte
	indexSettings strings.Builder

	samples    []float64
	throughput map[string]float64
}

type Config struct {
	IndexName   string
	DatasetName string
	Client      *elasticsearch.Client
	Decoder     esutil.BulkResponseJSONDecoder

	NumShards     int
	NumReplicas   int
	NumItems      int
	NumRuns       int
	NumWarmupRuns int
	NumWorkers    int
	FlushBytes    int
	Wait          time.Duration
	Mockserver    bool
}

func (r *Runner) Report() error { _ = "STUB: not implemented"; return nil }

func (r *Runner) Run() error { _ = "STUB: not implemented"; return nil }

func (r *Runner) setup() error { _ = "STUB: not implemented"; return nil }

func (r *Runner) run(n int, measure bool) error { _ = "STUB: not implemented"; return nil }

func formatInt(i int) string { _ = "STUB: not implemented"; return "" }
