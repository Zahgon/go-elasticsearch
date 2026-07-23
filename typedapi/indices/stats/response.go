package stats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	All_    types.IndicesStats            `json:"_all"`
	Indices map[string]types.IndicesStats `json:"indices,omitempty"`
	Shards_ types.ShardStatistics         `json:"_shards"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
