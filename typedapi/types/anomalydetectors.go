package types

type AnomalyDetectors struct {
	CategorizationAnalyzer               CategorizationAnalyzer `json:"categorization_analyzer"`
	CategorizationExamplesLimit          int                    `json:"categorization_examples_limit"`
	DailyModelSnapshotRetentionAfterDays int                    `json:"daily_model_snapshot_retention_after_days"`
	ModelMemoryLimit                     string                 `json:"model_memory_limit"`
	ModelSnapshotRetentionDays           int                    `json:"model_snapshot_retention_days"`
}

func (s *AnomalyDetectors) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAnomalyDetectors() *AnomalyDetectors { _ = "STUB: not implemented"; return nil }
