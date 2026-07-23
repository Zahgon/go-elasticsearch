package updatebyquery

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Batches *int64 `json:"batches,omitempty"`

	Deleted *int64 `json:"deleted,omitempty"`

	Failures []types.BulkIndexByScrollFailure `json:"failures,omitempty"`

	Noops *int64 `json:"noops,omitempty"`

	RequestsPerSecond *float32 `json:"requests_per_second,omitempty"`

	Retries *types.Retries `json:"retries,omitempty"`

	Slices    []types.ReindexStatus `json:"slices,omitempty"`
	Task      *string               `json:"task,omitempty"`
	Throttled types.Duration        `json:"throttled,omitempty"`

	ThrottledMillis *int64         `json:"throttled_millis,omitempty"`
	ThrottledUntil  types.Duration `json:"throttled_until,omitempty"`

	ThrottledUntilMillis *int64 `json:"throttled_until_millis,omitempty"`

	TimedOut *bool `json:"timed_out,omitempty"`

	Took *int64 `json:"took,omitempty"`

	Total *int64 `json:"total,omitempty"`

	Updated *int64 `json:"updated,omitempty"`

	VersionConflicts *int64 `json:"version_conflicts,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
