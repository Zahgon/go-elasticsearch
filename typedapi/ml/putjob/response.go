package putjob

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	AllowLazyOpen                        bool                     `json:"allow_lazy_open"`
	AnalysisConfig                       types.AnalysisConfigRead `json:"analysis_config"`
	AnalysisLimits                       types.AnalysisLimits     `json:"analysis_limits"`
	BackgroundPersistInterval            types.Duration           `json:"background_persist_interval,omitempty"`
	CreateTime                           types.DateTime           `json:"create_time"`
	CustomSettings                       json.RawMessage          `json:"custom_settings,omitempty"`
	DailyModelSnapshotRetentionAfterDays int64                    `json:"daily_model_snapshot_retention_after_days"`
	DataDescription                      types.DataDescription    `json:"data_description"`
	DatafeedConfig                       *types.MLDatafeed        `json:"datafeed_config,omitempty"`
	Description                          *string                  `json:"description,omitempty"`
	Groups                               []string                 `json:"groups,omitempty"`
	JobId                                string                   `json:"job_id"`
	JobType                              string                   `json:"job_type"`
	JobVersion                           string                   `json:"job_version"`
	ModelPlotConfig                      *types.ModelPlotConfig   `json:"model_plot_config,omitempty"`
	ModelSnapshotId                      *string                  `json:"model_snapshot_id,omitempty"`
	ModelSnapshotRetentionDays           int64                    `json:"model_snapshot_retention_days"`
	RenormalizationWindowDays            *int64                   `json:"renormalization_window_days,omitempty"`
	ResultsIndexName                     string                   `json:"results_index_name"`
	ResultsRetentionDays                 *int64                   `json:"results_retention_days,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
