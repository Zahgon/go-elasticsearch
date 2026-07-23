package types

type HealthStatistics struct {
	Available   bool        `json:"available"`
	Enabled     bool        `json:"enabled"`
	Invocations Invocations `json:"invocations"`
}

func (s *HealthStatistics) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHealthStatistics() *HealthStatistics { _ = "STUB: not implemented"; return nil }
