package types

type StepChange struct {
	ChangePoint int     `json:"change_point"`
	PValue      Float64 `json:"p_value"`
}

func (s *StepChange) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewStepChange() *StepChange { _ = "STUB: not implemented"; return nil }
