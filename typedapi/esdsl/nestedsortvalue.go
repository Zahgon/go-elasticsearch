package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _nestedSortValue struct {
	v *types.NestedSortValue
}

func NewNestedSortValue() *_nestedSortValue { _ = "STUB: not implemented"; return nil }

func (s *_nestedSortValue) Filter(filter types.QueryVariant) *_nestedSortValue {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedSortValue) MaxChildren(maxchildren int) *_nestedSortValue {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedSortValue) Nested(nested types.NestedSortValueVariant) *_nestedSortValue {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedSortValue) Path(field string) *_nestedSortValue {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedSortValue) NestedSortValueCaster() *types.NestedSortValue {
	_ = "STUB: not implemented"
	return nil
}
