package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jobstate"
)

type JobStats struct {
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`

	DataCounts DataCounts `json:"data_counts"`

	Deleting *bool `json:"deleting,omitempty"`

	ForecastsStats JobForecastStatistics `json:"forecasts_stats"`

	JobId string `json:"job_id"`

	ModelSizeStats ModelSizeStats `json:"model_size_stats"`

	Node *DiscoveryNodeCompact `json:"node,omitempty"`

	OpenTime DateTime `json:"open_time,omitempty"`

	State jobstate.JobState `json:"state"`

	TimingStats JobTimingStats `json:"timing_stats"`
}

func (s *JobStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJobStats() *JobStats { _ = "STUB: not implemented"; return nil }
