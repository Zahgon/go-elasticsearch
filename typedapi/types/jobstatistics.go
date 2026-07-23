package types

type JobStatistics struct {
	Avg   Float64 `json:"avg"`
	Max   Float64 `json:"max"`
	Min   Float64 `json:"min"`
	Total Float64 `json:"total"`
}

func (s *JobStatistics) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJobStatistics() *JobStatistics { _ = "STUB: not implemented"; return nil }
