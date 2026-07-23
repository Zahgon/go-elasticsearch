package bulkdeleterole

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Deleted []string `json:"deleted,omitempty"`

	Errors *types.BulkError `json:"errors,omitempty"`

	NotFound []string `json:"not_found,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
