package putalias

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Filter *types.Query `json:"filter,omitempty"`

	IndexRouting *string `json:"index_routing,omitempty"`

	IsWriteIndex *bool `json:"is_write_index,omitempty"`

	Routing *string `json:"routing,omitempty"`

	SearchRouting *string `json:"search_routing,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
