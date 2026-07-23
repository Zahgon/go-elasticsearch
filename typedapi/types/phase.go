package types

type Phase struct {
	Actions *IlmActions `json:"actions,omitempty"`
	MinAge  Duration    `json:"min_age,omitempty"`
}

func (s *Phase) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPhase() *Phase { _ = "STUB: not implemented"; return nil }

type PhaseVariant interface {
	PhaseCaster() *Phase
}

func (s *Phase) PhaseCaster() *Phase { _ = "STUB: not implemented"; return nil }
