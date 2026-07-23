package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _uppercaseTokenFilter struct {
	v *types.UppercaseTokenFilter
}

func NewUppercaseTokenFilter() *_uppercaseTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_uppercaseTokenFilter) Version(versionstring string) *_uppercaseTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uppercaseTokenFilter) UppercaseTokenFilterCaster() *types.UppercaseTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
