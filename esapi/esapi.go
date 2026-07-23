package esapi

import (
	"net/http"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/internal/version"
)

const Version = version.Client

type Transport interface {
	Perform(*http.Request) (*http.Response, error)
}

type Instrumented elastictransport.Instrumented

type Instrumentation elastictransport.Instrumentation

func BoolPtr(v bool) *bool { _ = "STUB: not implemented"; return nil }

func IntPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func Int64Ptr(v int64) *int64 { _ = "STUB: not implemented"; return nil }

func formatDuration(d time.Duration) string { _ = "STUB: not implemented"; return "" }
