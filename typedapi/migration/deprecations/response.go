package deprecations

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ClusterSettings []types.Deprecation            `json:"cluster_settings"`
	DataStreams     map[string][]types.Deprecation `json:"data_streams"`

	IlmPolicies map[string][]types.Deprecation `json:"ilm_policies"`

	IndexSettings map[string][]types.Deprecation `json:"index_settings"`

	MlSettings []types.Deprecation `json:"ml_settings"`

	NodeSettings []types.Deprecation `json:"node_settings"`

	Templates map[string][]types.Deprecation `json:"templates"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
