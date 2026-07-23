package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanOrQuery struct {
	v *types.SpanOrQuery
}

func NewSpanOrQuery() *_spanOrQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanOrQuery) Clauses(clauses ...types.SpanQueryVariant) *_spanOrQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanOrQuery) ClausesValues(clausesvalues []types.SpanQuery) *_spanOrQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanOrQuery) Boost(boost float32) *_spanOrQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanOrQuery) QueryName_(queryname_ string) *_spanOrQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanOrQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_spanOrQuery) SpanQueryCaster() *types.SpanQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanOrQuery) SpanOrQueryCaster() *types.SpanOrQuery {
	_ = "STUB: not implemented"
	return nil
}
