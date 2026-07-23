package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icunormalizationmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icunormalizationtype"
)

type _icuNormalizationCharFilter struct {
	v *types.IcuNormalizationCharFilter
}

func NewIcuNormalizationCharFilter() *_icuNormalizationCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuNormalizationCharFilter) Mode(mode icunormalizationmode.IcuNormalizationMode) *_icuNormalizationCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuNormalizationCharFilter) Name(name icunormalizationtype.IcuNormalizationType) *_icuNormalizationCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuNormalizationCharFilter) UnicodeSetFilter(unicodesetfilter string) *_icuNormalizationCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuNormalizationCharFilter) Version(versionstring string) *_icuNormalizationCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuNormalizationCharFilter) IcuNormalizationCharFilterCaster() *types.IcuNormalizationCharFilter {
	_ = "STUB: not implemented"
	return nil
}
