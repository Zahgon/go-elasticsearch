package rollupsearch

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Aggregations map[string]types.Aggregations `json:"aggregations,omitempty"`

	Query *types.Query `json:"query,omitempty"`

	Size *int `json:"size,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
