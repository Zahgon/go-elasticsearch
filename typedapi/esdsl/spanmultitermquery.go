package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanMultiTermQuery struct {
	v *types.SpanMultiTermQuery
}

func NewSpanMultiTermQuery(match types.QueryVariant) *_spanMultiTermQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanMultiTermQuery) Match(match types.QueryVariant) *_spanMultiTermQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanMultiTermQuery) Boost(boost float32) *_spanMultiTermQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanMultiTermQuery) QueryName_(queryname_ string) *_spanMultiTermQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanMultiTermQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_spanMultiTermQuery) SpanQueryCaster() *types.SpanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanMultiTermQuery) SpanMultiTermQueryCaster() *types.SpanMultiTermQuery {
	_ = "STUB: not implemented"
	return nil
}
