package types

type SpanFieldMaskingQuery struct {
	Boost      *float32  `json:"boost,omitempty"`
	Field      string    `json:"field"`
	Query      SpanQuery `json:"query"`
	QueryName_ *string   `json:"_name,omitempty"`
}

func (s *SpanFieldMaskingQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSpanFieldMaskingQuery() *SpanFieldMaskingQuery { _ = "STUB: not implemented"; return nil }

type SpanFieldMaskingQueryVariant interface {
	SpanFieldMaskingQueryCaster() *SpanFieldMaskingQuery
}

func (s *SpanFieldMaskingQuery) SpanFieldMaskingQueryCaster() *SpanFieldMaskingQuery {
	_ = "STUB: not implemented"
	return nil
}
