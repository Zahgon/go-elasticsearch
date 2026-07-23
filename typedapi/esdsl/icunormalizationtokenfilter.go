package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icunormalizationtype"
)

type _icuNormalizationTokenFilter struct {
	v *types.IcuNormalizationTokenFilter
}

func NewIcuNormalizationTokenFilter(name icunormalizationtype.IcuNormalizationType) *_icuNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuNormalizationTokenFilter) Name(name icunormalizationtype.IcuNormalizationType) *_icuNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuNormalizationTokenFilter) Version(versionstring string) *_icuNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuNormalizationTokenFilter) IcuNormalizationTokenFilterCaster() *types.IcuNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
