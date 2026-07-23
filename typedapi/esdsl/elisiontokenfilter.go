package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _elisionTokenFilter struct {
	v *types.ElisionTokenFilter
}

func NewElisionTokenFilter() *_elisionTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_elisionTokenFilter) Articles(articles ...string) *_elisionTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elisionTokenFilter) ArticlesCase(stringifiedboolean types.StringifiedbooleanVariant) *_elisionTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elisionTokenFilter) ArticlesPath(articlespath string) *_elisionTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elisionTokenFilter) Version(versionstring string) *_elisionTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elisionTokenFilter) ElisionTokenFilterCaster() *types.ElisionTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
