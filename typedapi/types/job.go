package types

import (
	"encoding/json"
)

type Job struct {
	AllowLazyOpen bool `json:"allow_lazy_open"`

	AnalysisConfig AnalysisConfig `json:"analysis_config"`

	AnalysisLimits *AnalysisLimits `json:"analysis_limits,omitempty"`

	BackgroundPersistInterval Duration    `json:"background_persist_interval,omitempty"`
	Blocked                   *JobBlocked `json:"blocked,omitempty"`
	CreateTime                DateTime    `json:"create_time,omitempty"`

	CustomSettings json.RawMessage `json:"custom_settings,omitempty"`

	DailyModelSnapshotRetentionAfterDays *int64 `json:"daily_model_snapshot_retention_after_days,omitempty"`

	DataDescription DataDescription `json:"data_description"`

	DatafeedConfig *MLDatafeed `json:"datafeed_config,omitempty"`

	Deleting *bool `json:"deleting,omitempty"`

	Description *string `json:"description,omitempty"`

	FinishedTime DateTime `json:"finished_time,omitempty"`

	Groups []string `json:"groups,omitempty"`

	JobId string `json:"job_id"`

	JobType *string `json:"job_type,omitempty"`

	JobVersion *string `json:"job_version,omitempty"`

	ModelPlotConfig *ModelPlotConfig `json:"model_plot_config,omitempty"`
	ModelSnapshotId *string          `json:"model_snapshot_id,omitempty"`

	ModelSnapshotRetentionDays int64 `json:"model_snapshot_retention_days"`

	RenormalizationWindowDays *int64 `json:"renormalization_window_days,omitempty"`

	ResultsIndexName string `json:"results_index_name"`

	ResultsRetentionDays *int64 `json:"results_retention_days,omitempty"`
}

func (s *Job) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJob() *Job { _ = "STUB: not implemented"; return nil }
