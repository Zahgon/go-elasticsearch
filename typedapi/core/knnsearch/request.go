package knnsearch

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	DocvalueFields []types.FieldAndFormat `json:"docvalue_fields,omitempty"`

	Fields []string `json:"fields,omitempty"`

	Filter []types.Query `json:"filter,omitempty"`

	Knn types.KnnSearchQuery `json:"knn"`

	Source_ types.SourceConfig `json:"_source,omitempty"`

	StoredFields []string `json:"stored_fields,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
