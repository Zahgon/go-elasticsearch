package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanNotQuery struct {
	v *types.SpanNotQuery
}

func NewSpanNotQuery(exclude types.SpanQueryVariant, include types.SpanQueryVariant) *_spanNotQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanNotQuery) Dist(dist int) *_spanNotQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanNotQuery) Exclude(exclude types.SpanQueryVariant) *_spanNotQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanNotQuery) Include(include types.SpanQueryVariant) *_spanNotQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanNotQuery) Post(post int) *_spanNotQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanNotQuery) Pre(pre int) *_spanNotQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanNotQuery) Boost(boost float32) *_spanNotQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanNotQuery) QueryName_(queryname_ string) *_spanNotQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanNotQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_spanNotQuery) SpanQueryCaster() *types.SpanQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanNotQuery) SpanNotQueryCaster() *types.SpanNotQuery {
	_ = "STUB: not implemented"
	return nil
}
