package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _uniqueTokenFilter struct {
	v *types.UniqueTokenFilter
}

func NewUniqueTokenFilter() *_uniqueTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_uniqueTokenFilter) OnlyOnSamePosition(onlyonsameposition bool) *_uniqueTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uniqueTokenFilter) Version(versionstring string) *_uniqueTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uniqueTokenFilter) UniqueTokenFilterCaster() *types.UniqueTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
