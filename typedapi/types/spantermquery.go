package types

type SpanTermQuery struct {
	Boost      *float32   `json:"boost,omitempty"`
	QueryName_ *string    `json:"_name,omitempty"`
	Value      FieldValue `json:"value"`
}

func (s *SpanTermQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSpanTermQuery() *SpanTermQuery { _ = "STUB: not implemented"; return nil }

type SpanTermQueryVariant interface {
	SpanTermQueryCaster() *SpanTermQuery
}

func (s *SpanTermQuery) SpanTermQueryCaster() *SpanTermQuery { _ = "STUB: not implemented"; return nil }
