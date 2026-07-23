package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _hunspellTokenFilter struct {
	v *types.HunspellTokenFilter
}

func NewHunspellTokenFilter(locale string) *_hunspellTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hunspellTokenFilter) Dedup(dedup bool) *_hunspellTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hunspellTokenFilter) Dictionary(dictionary string) *_hunspellTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hunspellTokenFilter) Locale(locale string) *_hunspellTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hunspellTokenFilter) LongestOnly(longestonly bool) *_hunspellTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hunspellTokenFilter) Version(versionstring string) *_hunspellTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hunspellTokenFilter) HunspellTokenFilterCaster() *types.HunspellTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
