package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icutransformdirection"
)

type _icuTransformTokenFilter struct {
	v *types.IcuTransformTokenFilter
}

func NewIcuTransformTokenFilter(id string) *_icuTransformTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuTransformTokenFilter) Dir(dir icutransformdirection.IcuTransformDirection) *_icuTransformTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuTransformTokenFilter) Id(id string) *_icuTransformTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuTransformTokenFilter) Version(versionstring string) *_icuTransformTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuTransformTokenFilter) IcuTransformTokenFilterCaster() *types.IcuTransformTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
