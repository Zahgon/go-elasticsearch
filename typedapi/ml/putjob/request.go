package putjob

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	AllowLazyOpen *bool `json:"allow_lazy_open,omitempty"`

	AnalysisConfig types.AnalysisConfig `json:"analysis_config"`

	AnalysisLimits *types.AnalysisLimits `json:"analysis_limits,omitempty"`

	BackgroundPersistInterval types.Duration `json:"background_persist_interval,omitempty"`

	CustomSettings json.RawMessage `json:"custom_settings,omitempty"`

	DailyModelSnapshotRetentionAfterDays *int64 `json:"daily_model_snapshot_retention_after_days,omitempty"`

	DataDescription types.DataDescription `json:"data_description"`

	DatafeedConfig *types.DatafeedConfig `json:"datafeed_config,omitempty"`

	Description *string `json:"description,omitempty"`

	Groups []string `json:"groups,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	ModelPlotConfig *types.ModelPlotConfig `json:"model_plot_config,omitempty"`

	ModelSnapshotRetentionDays *int64 `json:"model_snapshot_retention_days,omitempty"`

	RenormalizationWindowDays *int64 `json:"renormalization_window_days,omitempty"`

	ResultsIndexName *string `json:"results_index_name,omitempty"`

	ResultsRetentionDays *int64 `json:"results_retention_days,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
