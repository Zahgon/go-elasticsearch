package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/childscoremode"
)

type HasChildQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	IgnoreUnmapped *bool `json:"ignore_unmapped,omitempty"`

	InnerHits *InnerHits `json:"inner_hits,omitempty"`

	MaxChildren *int `json:"max_children,omitempty"`

	MinChildren *int `json:"min_children,omitempty"`

	Query      Query   `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`

	ScoreMode *childscoremode.ChildScoreMode `json:"score_mode,omitempty"`

	Type string `json:"type"`
}

func (s *HasChildQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHasChildQuery() *HasChildQuery { _ = "STUB: not implemented"; return nil }

type HasChildQueryVariant interface {
	HasChildQueryCaster() *HasChildQuery
}

func (s *HasChildQuery) HasChildQueryCaster() *HasChildQuery { _ = "STUB: not implemented"; return nil }
