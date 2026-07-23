package types

type JobUsage struct {
	Count     int              `json:"count"`
	CreatedBy map[string]int64 `json:"created_by"`
	Detectors JobStatistics    `json:"detectors"`
	Forecasts MlJobForecasts   `json:"forecasts"`
	ModelSize JobStatistics    `json:"model_size"`
}

func (s *JobUsage) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJobUsage() *JobUsage { _ = "STUB: not implemented"; return nil }
