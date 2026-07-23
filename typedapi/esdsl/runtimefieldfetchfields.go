package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _runtimeFieldFetchFields struct {
	v *types.RuntimeFieldFetchFields
}

func NewRuntimeFieldFetchFields() *_runtimeFieldFetchFields { _ = "STUB: not implemented"; return nil }

func (s *_runtimeFieldFetchFields) Field(field string) *_runtimeFieldFetchFields {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeFieldFetchFields) Format(format string) *_runtimeFieldFetchFields {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeFieldFetchFields) RuntimeFieldFetchFieldsCaster() *types.RuntimeFieldFetchFields {
	_ = "STUB: not implemented"
	return nil
}
