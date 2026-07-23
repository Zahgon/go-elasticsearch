package types

type SpanMultiTermQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Match      Query   `json:"match"`
	QueryName_ *string `json:"_name,omitempty"`
}

func (s *SpanMultiTermQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSpanMultiTermQuery() *SpanMultiTermQuery { _ = "STUB: not implemented"; return nil }

type SpanMultiTermQueryVariant interface {
	SpanMultiTermQueryCaster() *SpanMultiTermQuery
}

func (s *SpanMultiTermQuery) SpanMultiTermQueryCaster() *SpanMultiTermQuery {
	_ = "STUB: not implemented"
	return nil
}
