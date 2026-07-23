package reloadsecuresettings

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ClusterName string `json:"cluster_name"`

	NodeStats *types.NodeStatistics             `json:"_nodes,omitempty"`
	Nodes     map[string]types.NodeReloadResult `json:"nodes"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
