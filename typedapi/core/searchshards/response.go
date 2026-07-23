package searchshards

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Indices map[string]types.ShardStoreIndex            `json:"indices"`
	Nodes   map[string]types.SearchShardsNodeAttributes `json:"nodes"`
	Shards  [][]types.NodeShard                         `json:"shards"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
