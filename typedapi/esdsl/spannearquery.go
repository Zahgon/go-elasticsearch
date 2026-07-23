package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanNearQuery struct {
	v *types.SpanNearQuery
}

func NewSpanNearQuery() *_spanNearQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanNearQuery) Clauses(clauses ...types.SpanQueryVariant) *_spanNearQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanNearQuery) ClausesValues(clausesvalues []types.SpanQuery) *_spanNearQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanNearQuery) InOrder(inorder bool) *_spanNearQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanNearQuery) Slop(slop int) *_spanNearQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanNearQuery) Boost(boost float32) *_spanNearQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanNearQuery) QueryName_(queryname_ string) *_spanNearQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanNearQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_spanNearQuery) SpanQueryCaster() *types.SpanQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanNearQuery) SpanNearQueryCaster() *types.SpanNearQuery {
	_ = "STUB: not implemented"
	return nil
}
