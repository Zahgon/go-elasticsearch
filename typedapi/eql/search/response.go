package search

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Hits types.EqlHits `json:"hits"`

	Id *string `json:"id,omitempty"`

	IsPartial *bool `json:"is_partial,omitempty"`

	IsRunning *bool `json:"is_running,omitempty"`

	ShardFailures []types.ShardFailure `json:"shard_failures,omitempty"`

	TimedOut *bool `json:"timed_out,omitempty"`

	Took *int64 `json:"took,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
