package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanGapQuery struct {
	v types.SpanGapQuery
}

func NewSpanGapQuery(spangapquery map[string]int) *_spanGapQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_spanGapQuery) SpanGapQueryCaster() *types.SpanGapQuery {
	_ = "STUB: not implemented"
	return nil
}
