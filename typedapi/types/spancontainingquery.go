package types

type SpanContainingQuery struct {
	Big SpanQuery `json:"big"`

	Boost *float32 `json:"boost,omitempty"`

	Little     SpanQuery `json:"little"`
	QueryName_ *string   `json:"_name,omitempty"`
}

func (s *SpanContainingQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSpanContainingQuery() *SpanContainingQuery { _ = "STUB: not implemented"; return nil }

type SpanContainingQueryVariant interface {
	SpanContainingQueryCaster() *SpanContainingQuery
}

func (s *SpanContainingQuery) SpanContainingQueryCaster() *SpanContainingQuery {
	_ = "STUB: not implemented"
	return nil
}
