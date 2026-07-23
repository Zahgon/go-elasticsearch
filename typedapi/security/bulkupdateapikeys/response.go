package bulkupdateapikeys

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Errors  *types.BulkError `json:"errors,omitempty"`
	Noops   []string         `json:"noops"`
	Updated []string         `json:"updated"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
