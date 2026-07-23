package validate

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	AnalysisConfig             *types.AnalysisConfig  `json:"analysis_config,omitempty"`
	AnalysisLimits             *types.AnalysisLimits  `json:"analysis_limits,omitempty"`
	DataDescription            *types.DataDescription `json:"data_description,omitempty"`
	Description                *string                `json:"description,omitempty"`
	JobId                      *string                `json:"job_id,omitempty"`
	ModelPlot                  *types.ModelPlotConfig `json:"model_plot,omitempty"`
	ModelSnapshotId            *string                `json:"model_snapshot_id,omitempty"`
	ModelSnapshotRetentionDays *int64                 `json:"model_snapshot_retention_days,omitempty"`
	ResultsIndexName           *string                `json:"results_index_name,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
