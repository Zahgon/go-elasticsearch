package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _parentIdQuery struct {
	v *types.ParentIdQuery
}

func NewParentIdQuery() *_parentIdQuery { _ = "STUB: not implemented"; return nil }

func (s *_parentIdQuery) Id(id string) *_parentIdQuery { _ = "STUB: not implemented"; return nil }

func (s *_parentIdQuery) IgnoreUnmapped(ignoreunmapped bool) *_parentIdQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_parentIdQuery) Type(relationname string) *_parentIdQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_parentIdQuery) Boost(boost float32) *_parentIdQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_parentIdQuery) QueryName_(queryname_ string) *_parentIdQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_parentIdQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_parentIdQuery) ParentIdQueryCaster() *types.ParentIdQuery {
	_ = "STUB: not implemented"
	return nil
}
