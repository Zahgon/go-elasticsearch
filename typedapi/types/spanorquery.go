package types

type SpanOrQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Clauses    []SpanQuery `json:"clauses"`
	QueryName_ *string     `json:"_name,omitempty"`
}

func (s *SpanOrQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSpanOrQuery() *SpanOrQuery { _ = "STUB: not implemented"; return nil }

type SpanOrQueryVariant interface {
	SpanOrQueryCaster() *SpanOrQuery
}

func (s *SpanOrQuery) SpanOrQueryCaster() *SpanOrQuery { _ = "STUB: not implemented"; return nil }
