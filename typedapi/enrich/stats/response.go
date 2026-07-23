package stats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	CacheStats []types.CacheStats `json:"cache_stats,omitempty"`

	CoordinatorStats []types.CoordinatorStats `json:"coordinator_stats"`

	ExecutingPolicies []types.ExecutingPolicy `json:"executing_policies"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
