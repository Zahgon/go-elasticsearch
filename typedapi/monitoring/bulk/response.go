package bulk

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Error *types.ErrorCause `json:"error,omitempty"`

	Errors bool `json:"errors"`

	Ignored bool  `json:"ignored"`
	Took    int64 `json:"took"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
