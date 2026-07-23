package invalidateapikey

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ErrorCount int `json:"error_count"`

	ErrorDetails []types.ErrorCause `json:"error_details,omitempty"`

	InvalidatedApiKeys []string `json:"invalidated_api_keys"`

	PreviouslyInvalidatedApiKeys []string `json:"previously_invalidated_api_keys"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
