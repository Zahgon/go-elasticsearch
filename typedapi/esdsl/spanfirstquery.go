package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanFirstQuery struct {
	v *types.SpanFirstQuery
}

func NewSpanFirstQuery(end int, match types.SpanQueryVariant) *_spanFirstQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanFirstQuery) End(end int) *_spanFirstQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanFirstQuery) Match(match types.SpanQueryVariant) *_spanFirstQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanFirstQuery) Boost(boost float32) *_spanFirstQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanFirstQuery) QueryName_(queryname_ string) *_spanFirstQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanFirstQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_spanFirstQuery) SpanQueryCaster() *types.SpanQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanFirstQuery) SpanFirstQueryCaster() *types.SpanFirstQuery {
	_ = "STUB: not implemented"
	return nil
}
