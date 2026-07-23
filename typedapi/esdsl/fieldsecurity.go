package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fieldSecurity struct {
	v *types.FieldSecurity
}

func NewFieldSecurity() *_fieldSecurity { _ = "STUB: not implemented"; return nil }

func (s *_fieldSecurity) Except(fields ...string) *_fieldSecurity {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSecurity) Grant(fields ...string) *_fieldSecurity {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSecurity) FieldSecurityCaster() *types.FieldSecurity {
	_ = "STUB: not implemented"
	return nil
}
