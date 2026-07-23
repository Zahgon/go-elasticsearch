package runner

import (
	"regexp"
	"runtime"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
	"github.com/elastic/go-elasticsearch/v9/esutil"
)

var (
	statsIndex = "metrics-intake-" + time.Now().Format("2006-01")

	reGoVersion = regexp.MustCompile(`go(\d+\.\d+\.?.?)`)
	errs        []error

	RuntimeOS      string
	RuntimeVersion string
)

func init() {
	RuntimeOS = runtime.GOOS

	if v := reGoVersion.ReplaceAllString(runtime.Version(), "$1"); v != "" {
		RuntimeVersion = v
	} else {
		RuntimeVersion = runtime.Version()
	}
}

func NewRunner(cfg Config) (*Runner, error) { _ = "STUB: not implemented"; return nil, nil }

type Runner struct {
	config  Config
	stats   []Stats
	indexer esutil.BulkIndexer
}

type Config struct {
	BuildID string

	Action      string
	Category    string
	Environment string

	NumWarmups     int
	NumRepetitions int
	NumOperations  int

	SetupFunc    RunnerFunc
	RunnerFunc   RunnerFunc
	RunnerClient *elasticsearch.Client

	ReportClient *elasticsearch.Client

	Target struct {
		OS      ConfigOS
		Service ConfigService
	}

	Runner struct {
		Service ConfigService
	}
}

type ConfigOS struct {
	Family string
}

type ConfigService struct {
	Type    string
	Name    string
	Version string
	Git     ConfigGit
}

type ConfigGit struct {
	Branch string
	Commit string
}

type RunnerFunc func(int, Config) (*esapi.Response, error)

type Stats struct {
	Start              time.Time
	Duration           time.Duration
	Outcome            string
	ResponseStatusCode int
}

type Error struct {
	err  string
	errs []error
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) Errs() string { _ = "STUB: not implemented"; return "" }

func (r *Runner) Run() error { _ = "STUB: not implemented"; return nil }

func (r *Runner) Stats() []Stats { _ = "STUB: not implemented"; return nil }

func (r *Runner) Errs() []error { _ = "STUB: not implemented"; return nil }

func (r *Runner) SaveStats() error { _ = "STUB: not implemented"; return nil }

type record struct {
	Timestamp time.Time         `json:"@timestamp"`
	Labels    map[string]string `json:"labels,omitempty"`
	Tags      []string          `json:"tags,omitempty"`

	Event     recordEvent     `json:"event"`
	HTTP      recordHTTP      `json:"http,omitempty"`
	Benchmark recordBenchmark `json:"benchmark"`
}

type recordEvent struct {
	Action   string `json:"action"`
	Duration int64  `json:"duration"`
	Outcome  string `json:"outcome,omitempty"`
	Dataset  string `json:"dataset,omitempty"`
}

type recordBenchmark struct {
	BuildID     string       `json:"build_id"`
	Repetitions int          `json:"repetitions"`
	Operations  int          `json:"operations"`
	Runner      recordRunner `json:"runner"`
	Target      recordTarget `json:"target"`
	Category    string       `json:"category,omitempty"`
	Environment string       `json:"environment,omitempty"`
}

type recordRunner struct {
	Service recordService `json:"service"`
	Runtime recordRuntime `json:"runtime"`
	OS      recordOS      `json:"os"`
}

type recordTarget struct {
	Service recordService `json:"service"`
	OS      recordOS      `json:"os"`
}

type recordService struct {
	Type    string    `json:"type"`
	Name    string    `json:"name"`
	Version string    `json:"version"`
	Git     recordGit `json:"git,omitempty"`
}

type recordHTTP struct {
	Response recordHTTPResponse `json:"response,omitempty"`
}

type recordHTTPResponse struct {
	StatusCode int `json:"status_code"`
}

type recordOS struct {
	Family string `json:"family"`
}

type recordRuntime struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type recordGit struct {
	Branch string `json:"branch,omitempty"`
	Commit string `json:"commit,omitempty"`
}

func validateConfig(cfg Config) error { _ = "STUB: not implemented"; return nil }
