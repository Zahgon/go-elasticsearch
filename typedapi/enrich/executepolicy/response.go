package executepolicy

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Status *types.ExecuteEnrichPolicyStatus `json:"status,omitempty"`
	Task   *string                          `json:"task,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
