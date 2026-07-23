package bulkputrole

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Created []string `json:"created,omitempty"`

	Errors *types.BulkError `json:"errors,omitempty"`

	Noop []string `json:"noop,omitempty"`

	Updated []string `json:"updated,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
