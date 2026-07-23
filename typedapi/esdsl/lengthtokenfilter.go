package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _lengthTokenFilter struct {
	v *types.LengthTokenFilter
}

func NewLengthTokenFilter() *_lengthTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_lengthTokenFilter) Max(max int) *_lengthTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lengthTokenFilter) Min(min int) *_lengthTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lengthTokenFilter) Version(versionstring string) *_lengthTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lengthTokenFilter) LengthTokenFilterCaster() *types.LengthTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
