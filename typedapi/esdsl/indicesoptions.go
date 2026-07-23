package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

type _indicesOptions struct {
	v *types.IndicesOptions
}

func NewIndicesOptions() *_indicesOptions { _ = "STUB: not implemented"; return nil }

func (s *_indicesOptions) AllowNoIndices(allownoindices bool) *_indicesOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesOptions) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *_indicesOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesOptions) IgnoreThrottled(ignorethrottled bool) *_indicesOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesOptions) IgnoreUnavailable(ignoreunavailable bool) *_indicesOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesOptions) IndicesOptionsCaster() *types.IndicesOptions {
	_ = "STUB: not implemented"
	return nil
}
