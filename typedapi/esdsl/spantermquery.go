package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanTermQuery struct {
	k string
	v *types.SpanTermQuery
}

func NewSpanTermQuery(field string, value types.FieldValueVariant) *_spanTermQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanTermQuery) Value(fieldvalue types.FieldValueVariant) *_spanTermQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanTermQuery) Boost(boost float32) *_spanTermQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanTermQuery) QueryName_(queryname_ string) *_spanTermQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanTermQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_spanTermQuery) SpanQueryCaster() *types.SpanQuery { _ = "STUB: not implemented"; return nil }

func NewSingleSpanTermQuery() *_spanTermQuery { _ = "STUB: not implemented"; return nil }

func (s *_spanTermQuery) SpanTermQueryCaster() *types.SpanTermQuery {
	_ = "STUB: not implemented"
	return nil
}
