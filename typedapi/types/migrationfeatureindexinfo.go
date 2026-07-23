package types

type MigrationFeatureIndexInfo struct {
	FailureCause *ErrorCause `json:"failure_cause,omitempty"`
	Index        string      `json:"index"`
	Version      string      `json:"version"`
}

func (s *MigrationFeatureIndexInfo) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMigrationFeatureIndexInfo() *MigrationFeatureIndexInfo {
	_ = "STUB: not implemented"
	return nil
}
