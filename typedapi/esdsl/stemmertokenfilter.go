package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _stemmerTokenFilter struct {
	v *types.StemmerTokenFilter
}

func NewStemmerTokenFilter() *_stemmerTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_stemmerTokenFilter) Language(language string) *_stemmerTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stemmerTokenFilter) Version(versionstring string) *_stemmerTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stemmerTokenFilter) StemmerTokenFilterCaster() *types.StemmerTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
