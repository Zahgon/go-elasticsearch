package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sourceFilter struct {
	v *types.SourceFilter
}

func NewSourceFilter() *_sourceFilter { _ = "STUB: not implemented"; return nil }

func (s *_sourceFilter) ExcludeVectors(excludevectors bool) *_sourceFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sourceFilter) Excludes(fields ...string) *_sourceFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sourceFilter) Includes(fields ...string) *_sourceFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sourceFilter) SourceFilterCaster() *types.SourceFilter {
	_ = "STUB: not implemented"
	return nil
}
