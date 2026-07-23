package types

type ParentIdQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Id *string `json:"id,omitempty"`

	IgnoreUnmapped *bool   `json:"ignore_unmapped,omitempty"`
	QueryName_     *string `json:"_name,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (s *ParentIdQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewParentIdQuery() *ParentIdQuery { _ = "STUB: not implemented"; return nil }

type ParentIdQueryVariant interface {
	ParentIdQueryCaster() *ParentIdQuery
}

func (s *ParentIdQuery) ParentIdQueryCaster() *ParentIdQuery { _ = "STUB: not implemented"; return nil }
