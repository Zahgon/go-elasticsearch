package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _kStemTokenFilter struct {
	v *types.KStemTokenFilter
}

func NewKStemTokenFilter() *_kStemTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_kStemTokenFilter) Version(versionstring string) *_kStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_kStemTokenFilter) KStemTokenFilterCaster() *types.KStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
