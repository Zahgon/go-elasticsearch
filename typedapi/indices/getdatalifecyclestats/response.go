package getdatalifecyclestats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	DataStreamCount int `json:"data_stream_count"`

	DataStreams []types.DataStreamStats `json:"data_streams"`

	LastRunDurationInMillis *int64 `json:"last_run_duration_in_millis,omitempty"`

	TimeBetweenStartsInMillis *int64 `json:"time_between_starts_in_millis,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
