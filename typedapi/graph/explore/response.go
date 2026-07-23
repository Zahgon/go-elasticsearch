package explore

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Connections []types.Connection   `json:"connections"`
	Failures    []types.ShardFailure `json:"failures"`
	TimedOut    bool                 `json:"timed_out"`
	Took        int64                `json:"took"`
	Vertices    []types.Vertex       `json:"vertices"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
