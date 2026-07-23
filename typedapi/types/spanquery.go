package types

import (
	"encoding/json"
)

type SpanQuery struct {
	AdditionalSpanQueryProperty map[string]json.RawMessage `json:"-"`

	SpanContaining *SpanContainingQuery `json:"span_containing,omitempty"`

	SpanFieldMasking *SpanFieldMaskingQuery `json:"span_field_masking,omitempty"`

	SpanFirst *SpanFirstQuery `json:"span_first,omitempty"`
	SpanGap   SpanGapQuery    `json:"span_gap,omitempty"`

	SpanMulti *SpanMultiTermQuery `json:"span_multi,omitempty"`

	SpanNear *SpanNearQuery `json:"span_near,omitempty"`

	SpanNot *SpanNotQuery `json:"span_not,omitempty"`

	SpanOr *SpanOrQuery `json:"span_or,omitempty"`

	SpanTerm map[string]SpanTermQuery `json:"span_term,omitempty"`

	SpanWithin *SpanWithinQuery `json:"span_within,omitempty"`
}

func (s *SpanQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s SpanQuery) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewSpanQuery() *SpanQuery { _ = "STUB: not implemented"; return nil }

type SpanQueryVariant interface {
	SpanQueryCaster() *SpanQuery
}

func (s *SpanQuery) SpanQueryCaster() *SpanQuery { _ = "STUB: not implemented"; return nil }
