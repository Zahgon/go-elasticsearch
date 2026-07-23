package updatejob

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	AllowLazyOpen  *bool                      `json:"allow_lazy_open,omitempty"`
	AnalysisLimits *types.AnalysisMemoryLimit `json:"analysis_limits,omitempty"`

	BackgroundPersistInterval types.Duration `json:"background_persist_interval,omitempty"`
	CategorizationFilters     []string       `json:"categorization_filters,omitempty"`

	CustomSettings map[string]json.RawMessage `json:"custom_settings,omitempty"`

	DailyModelSnapshotRetentionAfterDays *int64 `json:"daily_model_snapshot_retention_after_days,omitempty"`

	Description *string `json:"description,omitempty"`

	Detectors []types.DetectorUpdate `json:"detectors,omitempty"`

	Groups           []string               `json:"groups,omitempty"`
	ModelPlotConfig  *types.ModelPlotConfig `json:"model_plot_config,omitempty"`
	ModelPruneWindow types.Duration         `json:"model_prune_window,omitempty"`

	ModelSnapshotRetentionDays *int64 `json:"model_snapshot_retention_days,omitempty"`

	PerPartitionCategorization *types.PerPartitionCategorization `json:"per_partition_categorization,omitempty"`

	RenormalizationWindowDays *int64 `json:"renormalization_window_days,omitempty"`

	ResultsRetentionDays *int64 `json:"results_retention_days,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
