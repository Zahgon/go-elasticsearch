package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _keepWordsTokenFilter struct {
	v *types.KeepWordsTokenFilter
}

func NewKeepWordsTokenFilter() *_keepWordsTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_keepWordsTokenFilter) KeepWords(keepwords ...string) *_keepWordsTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keepWordsTokenFilter) KeepWordsCase(keepwordscase bool) *_keepWordsTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keepWordsTokenFilter) KeepWordsPath(keepwordspath string) *_keepWordsTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keepWordsTokenFilter) Version(versionstring string) *_keepWordsTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keepWordsTokenFilter) KeepWordsTokenFilterCaster() *types.KeepWordsTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
