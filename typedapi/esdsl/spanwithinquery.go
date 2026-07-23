package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanWithinQuery struct {
	v *types.SpanWithinQuery
}

func NewSpanWithinQuery(big types.SpanQueryVariant, little types.SpanQueryVariant) *_spanWithinQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanWithinQuery) Big(big types.SpanQueryVariant) *_spanWithinQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanWithinQuery) Little(little types.SpanQueryVariant) *_spanWithinQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanWithinQuery) Boost(boost float32) *_spanWithinQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanWithinQuery) QueryName_(queryname_ string) *_spanWithinQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanWithinQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_spanWithinQuery) SpanQueryCaster() *types.SpanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanWithinQuery) SpanWithinQueryCaster() *types.SpanWithinQuery {
	_ = "STUB: not implemented"
	return nil
}
