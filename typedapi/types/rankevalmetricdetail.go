package types

import (
	"encoding/json"
)

type RankEvalMetricDetail struct {
	Hits []RankEvalHitItem `json:"hits"`

	MetricDetails map[string]map[string]json.RawMessage `json:"metric_details"`

	MetricScore Float64 `json:"metric_score"`

	UnratedDocs []UnratedDocument `json:"unrated_docs"`
}

func (s *RankEvalMetricDetail) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRankEvalMetricDetail() *RankEvalMetricDetail { _ = "STUB: not implemented"; return nil }
