package info

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ClusterName string                       `json:"cluster_name"`
	Http        *types.Http                  `json:"http,omitempty"`
	Ingest      *types.NodesIngest           `json:"ingest,omitempty"`
	Script      *types.Scripting             `json:"script,omitempty"`
	ThreadPool  map[string]types.ThreadCount `json:"thread_pool,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
