package types

import (
	"encoding/json"
)

type JobConfig struct {
	AllowLazyOpen *bool `json:"allow_lazy_open,omitempty"`

	AnalysisConfig AnalysisConfig `json:"analysis_config"`

	AnalysisLimits *AnalysisLimits `json:"analysis_limits,omitempty"`

	BackgroundPersistInterval Duration `json:"background_persist_interval,omitempty"`

	CustomSettings json.RawMessage `json:"custom_settings,omitempty"`

	DailyModelSnapshotRetentionAfterDays *int64 `json:"daily_model_snapshot_retention_after_days,omitempty"`

	DataDescription DataDescription `json:"data_description"`

	DatafeedConfig *DatafeedConfig `json:"datafeed_config,omitempty"`

	Description *string `json:"description,omitempty"`

	Groups []string `json:"groups,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	JobType *string `json:"job_type,omitempty"`

	ModelPlotConfig *ModelPlotConfig `json:"model_plot_config,omitempty"`

	ModelSnapshotRetentionDays *int64 `json:"model_snapshot_retention_days,omitempty"`

	RenormalizationWindowDays *int64 `json:"renormalization_window_days,omitempty"`

	ResultsIndexName *string `json:"results_index_name,omitempty"`

	ResultsRetentionDays *int64 `json:"results_retention_days,omitempty"`
}

func (s *JobConfig) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJobConfig() *JobConfig { _ = "STUB: not implemented"; return nil }

type JobConfigVariant interface {
	JobConfigCaster() *JobConfig
}

func (s *JobConfig) JobConfigCaster() *JobConfig { _ = "STUB: not implemented"; return nil }
