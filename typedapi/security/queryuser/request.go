package queryuser

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	From *int `json:"from,omitempty"`

	Query *types.UserQueryContainer `json:"query,omitempty"`

	SearchAfter []types.FieldValue `json:"search_after,omitempty"`

	Size *int `json:"size,omitempty"`

	Sort []types.SortCombinations `json:"sort,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (r Request) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
