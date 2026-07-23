package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _jaStopTokenFilter struct {
	v *types.JaStopTokenFilter
}

func NewJaStopTokenFilter() *_jaStopTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_jaStopTokenFilter) Stopwords(stopwords types.StopWordsVariant) *_jaStopTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jaStopTokenFilter) Version(versionstring string) *_jaStopTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jaStopTokenFilter) JaStopTokenFilterCaster() *types.JaStopTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
