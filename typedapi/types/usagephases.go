package types

type UsagePhases struct {
	Cold   *UsagePhase `json:"cold,omitempty"`
	Delete *UsagePhase `json:"delete,omitempty"`
	Frozen *UsagePhase `json:"frozen,omitempty"`
	Hot    *UsagePhase `json:"hot,omitempty"`
	Warm   *UsagePhase `json:"warm,omitempty"`
}

func NewUsagePhases() *UsagePhases { _ = "STUB: not implemented"; return nil }
