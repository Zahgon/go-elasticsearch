package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _semanticQuery struct {
	v *types.SemanticQuery
}

func NewSemanticQuery(field string, query string) *_semanticQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticQuery) Field(field string) *_semanticQuery { _ = "STUB: not implemented"; return nil }

func (s *_semanticQuery) Query(query string) *_semanticQuery { _ = "STUB: not implemented"; return nil }

func (s *_semanticQuery) Boost(boost float32) *_semanticQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticQuery) QueryName_(queryname_ string) *_semanticQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_semanticQuery) SemanticQueryCaster() *types.SemanticQuery {
	_ = "STUB: not implemented"
	return nil
}
