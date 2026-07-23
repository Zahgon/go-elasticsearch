package stats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ClusterName *string `json:"cluster_name,omitempty"`

	NodeStats *types.NodeStatistics  `json:"_nodes,omitempty"`
	Nodes     map[string]types.Stats `json:"nodes"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
