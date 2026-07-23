package types

type DistributionChange struct {
	ChangePoint int     `json:"change_point"`
	PValue      Float64 `json:"p_value"`
}

func (s *DistributionChange) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDistributionChange() *DistributionChange { _ = "STUB: not implemented"; return nil }
