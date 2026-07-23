package types

type Phases struct {
	Cold   *Phase `json:"cold,omitempty"`
	Delete *Phase `json:"delete,omitempty"`
	Frozen *Phase `json:"frozen,omitempty"`
	Hot    *Phase `json:"hot,omitempty"`
	Warm   *Phase `json:"warm,omitempty"`
}

func NewPhases() *Phases { _ = "STUB: not implemented"; return nil }

type PhasesVariant interface {
	PhasesCaster() *Phases
}

func (s *Phases) PhasesCaster() *Phases { _ = "STUB: not implemented"; return nil }
