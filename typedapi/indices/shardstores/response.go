package shardstores

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Indices map[string]types.IndicesShardStores `json:"indices"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
