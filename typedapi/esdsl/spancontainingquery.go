package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanContainingQuery struct {
	v *types.SpanContainingQuery
}

func NewSpanContainingQuery(big types.SpanQueryVariant, little types.SpanQueryVariant) *_spanContainingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanContainingQuery) Big(big types.SpanQueryVariant) *_spanContainingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanContainingQuery) Little(little types.SpanQueryVariant) *_spanContainingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanContainingQuery) Boost(boost float32) *_spanContainingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanContainingQuery) QueryName_(queryname_ string) *_spanContainingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanContainingQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_spanContainingQuery) SpanQueryCaster() *types.SpanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanContainingQuery) SpanContainingQueryCaster() *types.SpanContainingQuery {
	_ = "STUB: not implemented"
	return nil
}
