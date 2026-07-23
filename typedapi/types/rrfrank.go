package types

type RrfRank struct {
	RankConstant *int64 `json:"rank_constant,omitempty"`

	RankWindowSize *int64 `json:"rank_window_size,omitempty"`
}

func (s *RrfRank) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRrfRank() *RrfRank { _ = "STUB: not implemented"; return nil }

type RrfRankVariant interface {
	RrfRankCaster() *RrfRank
}

func (s *RrfRank) RrfRankCaster() *RrfRank { _ = "STUB: not implemented"; return nil }
