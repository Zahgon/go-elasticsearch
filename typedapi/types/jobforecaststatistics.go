package types

type JobForecastStatistics struct {
	ForecastedJobs   int              `json:"forecasted_jobs"`
	MemoryBytes      *JobStatistics   `json:"memory_bytes,omitempty"`
	ProcessingTimeMs *JobStatistics   `json:"processing_time_ms,omitempty"`
	Records          *JobStatistics   `json:"records,omitempty"`
	Status           map[string]int64 `json:"status,omitempty"`
	Total            int64            `json:"total"`
}

func (s *JobForecastStatistics) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewJobForecastStatistics() *JobForecastStatistics { _ = "STUB: not implemented"; return nil }
