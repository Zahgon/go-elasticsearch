package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _stopTokenFilter struct {
	v *types.StopTokenFilter
}

func NewStopTokenFilter() *_stopTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_stopTokenFilter) IgnoreCase(ignorecase bool) *_stopTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stopTokenFilter) RemoveTrailing(removetrailing bool) *_stopTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stopTokenFilter) Stopwords(stopwords types.StopWordsVariant) *_stopTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stopTokenFilter) StopwordsPath(stopwordspath string) *_stopTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stopTokenFilter) Version(versionstring string) *_stopTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stopTokenFilter) StopTokenFilterCaster() *types.StopTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
