package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fieldNamesField struct {
	v *types.FieldNamesField
}

func NewFieldNamesField(enabled bool) *_fieldNamesField { _ = "STUB: not implemented"; return nil }

func (s *_fieldNamesField) Enabled(enabled bool) *_fieldNamesField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldNamesField) FieldNamesFieldCaster() *types.FieldNamesField {
	_ = "STUB: not implemented"
	return nil
}
