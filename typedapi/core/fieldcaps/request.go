package fieldcaps

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Fields []string `json:"fields,omitempty"`

	IndexFilter *types.Query `json:"index_filter,omitempty"`

	ProjectRouting *string `json:"project_routing,omitempty"`

	RuntimeMappings types.RuntimeFields `json:"runtime_mappings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
