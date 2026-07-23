package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/childscoremode"
)

type NestedQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	IgnoreUnmapped *bool `json:"ignore_unmapped,omitempty"`

	InnerHits *InnerHits `json:"inner_hits,omitempty"`

	Path string `json:"path"`

	Query      Query   `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`

	ScoreMode *childscoremode.ChildScoreMode `json:"score_mode,omitempty"`
}

func (s *NestedQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNestedQuery() *NestedQuery { _ = "STUB: not implemented"; return nil }

type NestedQueryVariant interface {
	NestedQueryCaster() *NestedQuery
}

func (s *NestedQuery) NestedQueryCaster() *NestedQuery { _ = "STUB: not implemented"; return nil }
