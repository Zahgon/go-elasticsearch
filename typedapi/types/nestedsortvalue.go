package types

type NestedSortValue struct {
	Filter      *Query           `json:"filter,omitempty"`
	MaxChildren *int             `json:"max_children,omitempty"`
	Nested      *NestedSortValue `json:"nested,omitempty"`
	Path        string           `json:"path"`
}

func (s *NestedSortValue) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNestedSortValue() *NestedSortValue { _ = "STUB: not implemented"; return nil }

type NestedSortValueVariant interface {
	NestedSortValueCaster() *NestedSortValue
}

func (s *NestedSortValue) NestedSortValueCaster() *NestedSortValue {
	_ = "STUB: not implemented"
	return nil
}
