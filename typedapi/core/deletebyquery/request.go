package deletebyquery

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	MaxDocs *int64 `json:"max_docs,omitempty"`

	Query *types.Query `json:"query,omitempty"`

	Slice *types.SlicedScroll `json:"slice,omitempty"`

	Sort []types.SortCombinations `json:"sort,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
