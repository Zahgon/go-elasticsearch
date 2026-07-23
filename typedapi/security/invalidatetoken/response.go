package invalidatetoken

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ErrorCount int64 `json:"error_count"`

	ErrorDetails []types.ErrorCause `json:"error_details,omitempty"`

	InvalidatedTokens int64 `json:"invalidated_tokens"`

	PreviouslyInvalidatedTokens int64 `json:"previously_invalidated_tokens"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
