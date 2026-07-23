package types

type MlJobForecasts struct {
	ForecastedJobs int64 `json:"forecasted_jobs"`
	Total          int64 `json:"total"`
}

func (s *MlJobForecasts) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMlJobForecasts() *MlJobForecasts { _ = "STUB: not implemented"; return nil }
