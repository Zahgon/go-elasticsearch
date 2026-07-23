package types

type LifecycleExplainPhaseExecution struct {
	ModifiedDateInMillis int64  `json:"modified_date_in_millis"`
	PhaseDefinition      *Phase `json:"phase_definition,omitempty"`
	Policy               string `json:"policy"`
	Version              int64  `json:"version"`
}

func (s *LifecycleExplainPhaseExecution) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewLifecycleExplainPhaseExecution() *LifecycleExplainPhaseExecution {
	_ = "STUB: not implemented"
	return nil
}
