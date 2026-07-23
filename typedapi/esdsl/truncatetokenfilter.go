package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _truncateTokenFilter struct {
	v *types.TruncateTokenFilter
}

func NewTruncateTokenFilter() *_truncateTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_truncateTokenFilter) Length(length int) *_truncateTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_truncateTokenFilter) Version(versionstring string) *_truncateTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_truncateTokenFilter) TruncateTokenFilterCaster() *types.TruncateTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
