package types

type IndexSettingsLifecycleStep struct {
	WaitTimeThreshold Duration `json:"wait_time_threshold,omitempty"`
}

func (s *IndexSettingsLifecycleStep) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexSettingsLifecycleStep() *IndexSettingsLifecycleStep {
	_ = "STUB: not implemented"
	return nil
}

type IndexSettingsLifecycleStepVariant interface {
	IndexSettingsLifecycleStepCaster() *IndexSettingsLifecycleStep
}

func (s *IndexSettingsLifecycleStep) IndexSettingsLifecycleStepCaster() *IndexSettingsLifecycleStep {
	_ = "STUB: not implemented"
	return nil
}
