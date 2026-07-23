package types

type TrendChange struct {
	ChangePoint int     `json:"change_point"`
	PValue      Float64 `json:"p_value"`
	RValue      Float64 `json:"r_value"`
}

func (s *TrendChange) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTrendChange() *TrendChange { _ = "STUB: not implemented"; return nil }
