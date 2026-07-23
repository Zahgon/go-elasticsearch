package info

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ClusterName string `json:"cluster_name"`

	NodeStats *types.NodeStatistics     `json:"_nodes,omitempty"`
	Nodes     map[string]types.NodeInfo `json:"nodes"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
