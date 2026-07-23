package validatequery

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Error        *string                              `json:"error,omitempty"`
	Explanations []types.IndicesValidationExplanation `json:"explanations,omitempty"`
	Shards_      *types.ShardStatistics               `json:"_shards,omitempty"`
	Valid        bool                                 `json:"valid"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
