package termsenum

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	CaseInsensitive *bool `json:"case_insensitive,omitempty"`

	Field string `json:"field"`

	IndexFilter *types.Query `json:"index_filter,omitempty"`

	ProjectRouting *string `json:"project_routing,omitempty"`

	SearchAfter *string `json:"search_after,omitempty"`

	Size *int `json:"size,omitempty"`

	String *string `json:"string,omitempty"`

	Timeout types.Duration `json:"timeout,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
