package types

import (
	"encoding/json"
)

type PercolateQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Document json.RawMessage `json:"document,omitempty"`

	Documents []json.RawMessage `json:"documents,omitempty"`

	Field string `json:"field"`

	Id *string `json:"id,omitempty"`

	Index *string `json:"index,omitempty"`

	Name *string `json:"name,omitempty"`

	Preference *string `json:"preference,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`

	Routing *string `json:"routing,omitempty"`

	Version *int64 `json:"version,omitempty"`
}

func (s *PercolateQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPercolateQuery() *PercolateQuery { _ = "STUB: not implemented"; return nil }

type PercolateQueryVariant interface {
	PercolateQueryCaster() *PercolateQuery
}

func (s *PercolateQuery) PercolateQueryCaster() *PercolateQuery {
	_ = "STUB: not implemented"
	return nil
}
