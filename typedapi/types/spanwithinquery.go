package types

type SpanWithinQuery struct {
	Big SpanQuery `json:"big"`

	Boost *float32 `json:"boost,omitempty"`

	Little     SpanQuery `json:"little"`
	QueryName_ *string   `json:"_name,omitempty"`
}

func (s *SpanWithinQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSpanWithinQuery() *SpanWithinQuery { _ = "STUB: not implemented"; return nil }

type SpanWithinQueryVariant interface {
	SpanWithinQueryCaster() *SpanWithinQuery
}

func (s *SpanWithinQuery) SpanWithinQueryCaster() *SpanWithinQuery {
	_ = "STUB: not implemented"
	return nil
}
