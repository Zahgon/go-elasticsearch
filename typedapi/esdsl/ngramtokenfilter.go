package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _nGramTokenFilter struct {
	v *types.NGramTokenFilter
}

func NewNGramTokenFilter() *_nGramTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_nGramTokenFilter) MaxGram(maxgram int) *_nGramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nGramTokenFilter) MinGram(mingram int) *_nGramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nGramTokenFilter) PreserveOriginal(stringifiedboolean types.StringifiedbooleanVariant) *_nGramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nGramTokenFilter) Version(versionstring string) *_nGramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nGramTokenFilter) NGramTokenFilterCaster() *types.NGramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
