package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scoremode"
)

type RescoreQuery struct {
	Query Query `json:"rescore_query"`

	QueryWeight *Float64 `json:"query_weight,omitempty"`

	RescoreQueryWeight *Float64 `json:"rescore_query_weight,omitempty"`

	ScoreMode *scoremode.ScoreMode `json:"score_mode,omitempty"`
}

func (s *RescoreQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRescoreQuery() *RescoreQuery { _ = "STUB: not implemented"; return nil }

type RescoreQueryVariant interface {
	RescoreQueryCaster() *RescoreQuery
}

func (s *RescoreQuery) RescoreQueryCaster() *RescoreQuery { _ = "STUB: not implemented"; return nil }
