package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _wrapperQuery struct {
	v *types.WrapperQuery
}

func NewWrapperQuery(query string) *_wrapperQuery { _ = "STUB: not implemented"; return nil }

func (s *_wrapperQuery) Query(query string) *_wrapperQuery { _ = "STUB: not implemented"; return nil }

func (s *_wrapperQuery) Boost(boost float32) *_wrapperQuery { _ = "STUB: not implemented"; return nil }

func (s *_wrapperQuery) QueryName_(queryname_ string) *_wrapperQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wrapperQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_wrapperQuery) WrapperQueryCaster() *types.WrapperQuery {
	_ = "STUB: not implemented"
	return nil
}
