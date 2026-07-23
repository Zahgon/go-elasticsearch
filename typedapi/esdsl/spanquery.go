package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _spanQuery struct {
	v *types.SpanQuery
}

func NewSpanQuery() *_spanQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanQuery) AdditionalSpanQueryProperty(key string, value json.RawMessage) *_spanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanQuery) SpanContaining(spancontaining types.SpanContainingQueryVariant) *_spanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanQuery) SpanFieldMasking(spanfieldmasking types.SpanFieldMaskingQueryVariant) *_spanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanQuery) SpanFirst(spanfirst types.SpanFirstQueryVariant) *_spanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanQuery) SpanGap(spangapquery types.SpanGapQueryVariant) *_spanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanQuery) SpanMulti(spanmulti types.SpanMultiTermQueryVariant) *_spanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanQuery) SpanNear(spannear types.SpanNearQueryVariant) *_spanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanQuery) SpanNot(spannot types.SpanNotQueryVariant) *_spanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanQuery) SpanOr(spanor types.SpanOrQueryVariant) *_spanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanQuery) SpanTerm(key string, value types.SpanTermQueryVariant) *_spanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanQuery) SpanWithin(spanwithin types.SpanWithinQueryVariant) *_spanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanQuery) SpanQueryCaster() *types.SpanQuery { _ = "STUB: not implemented"; return nil }
