package types

type RankEvalHit struct {
	Id_    string  `json:"_id"`
	Index_ string  `json:"_index"`
	Score_ Float64 `json:"_score"`
}

func (s *RankEvalHit) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRankEvalHit() *RankEvalHit { _ = "STUB: not implemented"; return nil }
