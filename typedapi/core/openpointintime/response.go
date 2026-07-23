package openpointintime

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Clusters_ *types.ClusterStatistics `json:"_clusters,omitempty"`
	Id        string                   `json:"id"`

	Shards_ types.ShardStatistics `json:"_shards"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
