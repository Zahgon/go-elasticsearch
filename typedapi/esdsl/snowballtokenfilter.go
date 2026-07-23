package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/snowballlanguage"
)

type _snowballTokenFilter struct {
	v *types.SnowballTokenFilter
}

func NewSnowballTokenFilter() *_snowballTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_snowballTokenFilter) Language(language snowballlanguage.SnowballLanguage) *_snowballTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_snowballTokenFilter) Version(versionstring string) *_snowballTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_snowballTokenFilter) SnowballTokenFilterCaster() *types.SnowballTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
