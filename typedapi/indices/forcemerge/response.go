package forcemerge

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Shards_ *types.ShardStatistics `json:"_shards,omitempty"`

	Task *string `json:"task,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
