package types

type Spike struct {
	ChangePoint int     `json:"change_point"`
	PValue      Float64 `json:"p_value"`
}

func (s *Spike) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSpike() *Spike { _ = "STUB: not implemented"; return nil }
