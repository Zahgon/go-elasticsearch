package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fieldCollapse struct {
	v *types.FieldCollapse
}

func NewFieldCollapse() *_fieldCollapse { _ = "STUB: not implemented"; return nil }

func (s *_fieldCollapse) Collapse(collapse types.FieldCollapseVariant) *_fieldCollapse {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldCollapse) Field(field string) *_fieldCollapse { _ = "STUB: not implemented"; return nil }

func (s *_fieldCollapse) InnerHits(innerhits ...types.InnerHitsVariant) *_fieldCollapse {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldCollapse) MaxConcurrentGroupSearches(maxconcurrentgroupsearches int) *_fieldCollapse {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldCollapse) FieldCollapseCaster() *types.FieldCollapse {
	_ = "STUB: not implemented"
	return nil
}
