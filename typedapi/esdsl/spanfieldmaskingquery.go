package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanFieldMaskingQuery struct {
	v *types.SpanFieldMaskingQuery
}

func NewSpanFieldMaskingQuery(query types.SpanQueryVariant) *_spanFieldMaskingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanFieldMaskingQuery) Field(field string) *_spanFieldMaskingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanFieldMaskingQuery) Query(query types.SpanQueryVariant) *_spanFieldMaskingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanFieldMaskingQuery) Boost(boost float32) *_spanFieldMaskingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanFieldMaskingQuery) QueryName_(queryname_ string) *_spanFieldMaskingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanFieldMaskingQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_spanFieldMaskingQuery) SpanQueryCaster() *types.SpanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanFieldMaskingQuery) SpanFieldMaskingQueryCaster() *types.SpanFieldMaskingQuery {
	_ = "STUB: not implemented"
	return nil
}
