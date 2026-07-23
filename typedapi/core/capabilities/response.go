package capabilities

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ClusterName string                      `json:"cluster_name"`
	Failures    []types.FailedNodeException `json:"failures,omitempty"`
	NodeStats   types.NodeStatistics        `json:"_nodes"`
	Supported   *bool                       `json:"supported,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
