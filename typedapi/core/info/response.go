package info

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ClusterName string `json:"cluster_name"`
	ClusterUuid string `json:"cluster_uuid"`

	Name    string `json:"name"`
	Tagline string `json:"tagline"`

	Version types.ElasticsearchVersionInfo `json:"version"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
