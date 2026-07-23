package types

type SpanNearQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Clauses []SpanQuery `json:"clauses"`

	InOrder    *bool   `json:"in_order,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`

	Slop *int `json:"slop,omitempty"`
}

func (s *SpanNearQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSpanNearQuery() *SpanNearQuery { _ = "STUB: not implemented"; return nil }

type SpanNearQueryVariant interface {
	SpanNearQueryCaster() *SpanNearQuery
}

func (s *SpanNearQuery) SpanNearQueryCaster() *SpanNearQuery { _ = "STUB: not implemented"; return nil }
