package types

type SpanNotQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Dist *int `json:"dist,omitempty"`

	Exclude SpanQuery `json:"exclude"`

	Include SpanQuery `json:"include"`

	Post *int `json:"post,omitempty"`

	Pre        *int    `json:"pre,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`
}

func (s *SpanNotQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSpanNotQuery() *SpanNotQuery { _ = "STUB: not implemented"; return nil }

type SpanNotQueryVariant interface {
	SpanNotQueryCaster() *SpanNotQuery
}

func (s *SpanNotQuery) SpanNotQueryCaster() *SpanNotQuery { _ = "STUB: not implemented"; return nil }
