package types

type QueryBreakdown struct {
	Advance                     int64 `json:"advance"`
	AdvanceCount                int64 `json:"advance_count"`
	BuildScorer                 int64 `json:"build_scorer"`
	BuildScorerCount            int64 `json:"build_scorer_count"`
	ComputeMaxScore             int64 `json:"compute_max_score"`
	ComputeMaxScoreCount        int64 `json:"compute_max_score_count"`
	CountWeight                 int64 `json:"count_weight"`
	CountWeightCount            int64 `json:"count_weight_count"`
	CreateWeight                int64 `json:"create_weight"`
	CreateWeightCount           int64 `json:"create_weight_count"`
	Match                       int64 `json:"match"`
	MatchCount                  int64 `json:"match_count"`
	NextDoc                     int64 `json:"next_doc"`
	NextDocCount                int64 `json:"next_doc_count"`
	Score                       int64 `json:"score"`
	ScoreCount                  int64 `json:"score_count"`
	SetMinCompetitiveScore      int64 `json:"set_min_competitive_score"`
	SetMinCompetitiveScoreCount int64 `json:"set_min_competitive_score_count"`
	ShallowAdvance              int64 `json:"shallow_advance"`
	ShallowAdvanceCount         int64 `json:"shallow_advance_count"`
}

func (s *QueryBreakdown) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewQueryBreakdown() *QueryBreakdown { _ = "STUB: not implemented"; return nil }
