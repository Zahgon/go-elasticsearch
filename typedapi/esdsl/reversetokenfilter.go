package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _reverseTokenFilter struct {
	v *types.ReverseTokenFilter
}

func NewReverseTokenFilter() *_reverseTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_reverseTokenFilter) Version(versionstring string) *_reverseTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reverseTokenFilter) ReverseTokenFilterCaster() *types.ReverseTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
