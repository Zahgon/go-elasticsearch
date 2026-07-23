package clearcachedservicetokens

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ClusterName string                       `json:"cluster_name"`
	NodeStats   types.NodeStatistics         `json:"_nodes"`
	Nodes       map[string]types.ClusterNode `json:"nodes"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
