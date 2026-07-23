package types

type RankEvalHitItem struct {
	Hit    RankEvalHit `json:"hit"`
	Rating *Float64    `json:"rating,omitempty"`
}

func (s *RankEvalHitItem) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRankEvalHitItem() *RankEvalHitItem { _ = "STUB: not implemented"; return nil }
