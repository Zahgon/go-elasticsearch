package benchmarks

import (
	"bytes"

	"github.com/elastic/go-elasticsearch/v9/benchmarks/runner"
)

var (
	Config      map[string]string
	Actions     []Action
	DataSources = make(map[string]*bytes.Buffer)

	DefaultRepetitions = 1000
)

type Action struct {
	Name           string
	Category       string
	Environment    string
	NumWarmups     int
	NumRepetitions int
	NumOperations  int
	SetupFunc      runner.RunnerFunc
	RunnerFunc     runner.RunnerFunc
}

func Register(a Action) error { _ = "STUB: not implemented"; return nil }
