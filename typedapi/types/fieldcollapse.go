package types

type FieldCollapse struct {
	Collapse *FieldCollapse `json:"collapse,omitempty"`

	Field string `json:"field"`

	InnerHits []InnerHits `json:"inner_hits,omitempty"`

	MaxConcurrentGroupSearches *int `json:"max_concurrent_group_searches,omitempty"`
}

func (s *FieldCollapse) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldCollapse() *FieldCollapse { _ = "STUB: not implemented"; return nil }

type FieldCollapseVariant interface {
	FieldCollapseCaster() *FieldCollapse
}

func (s *FieldCollapse) FieldCollapseCaster() *FieldCollapse { _ = "STUB: not implemented"; return nil }
