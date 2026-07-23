package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _trimTokenFilter struct {
	v *types.TrimTokenFilter
}

func NewTrimTokenFilter() *_trimTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_trimTokenFilter) Version(versionstring string) *_trimTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trimTokenFilter) TrimTokenFilterCaster() *types.TrimTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
