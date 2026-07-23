package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _typeQuery struct {
	v *types.TypeQuery
}

func NewTypeQuery(value string) *_typeQuery { _ = "STUB: not implemented"; return nil }

func (s *_typeQuery) Value(value string) *_typeQuery { _ = "STUB: not implemented"; return nil }

func (s *_typeQuery) Boost(boost float32) *_typeQuery { _ = "STUB: not implemented"; return nil }

func (s *_typeQuery) QueryName_(queryname_ string) *_typeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_typeQuery) TypeQueryCaster() *types.TypeQuery { _ = "STUB: not implemented"; return nil }
