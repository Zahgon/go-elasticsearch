package types

type SpanFirstQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	End int `json:"end"`

	Match      SpanQuery `json:"match"`
	QueryName_ *string   `json:"_name,omitempty"`
}

func (s *SpanFirstQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSpanFirstQuery() *SpanFirstQuery { _ = "STUB: not implemented"; return nil }

type SpanFirstQueryVariant interface {
	SpanFirstQueryCaster() *SpanFirstQuery
}

func (s *SpanFirstQuery) SpanFirstQueryCaster() *SpanFirstQuery {
	_ = "STUB: not implemented"
	return nil
}
