package count

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	ProjectRouting *string `json:"project_routing,omitempty"`

	Query *types.Query `json:"query,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
