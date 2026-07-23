package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _germanStemTokenFilter struct {
	v *types.GermanStemTokenFilter
}

func NewGermanStemTokenFilter() *_germanStemTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_germanStemTokenFilter) Version(versionstring string) *_germanStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_germanStemTokenFilter) GermanStemTokenFilterCaster() *types.GermanStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
