package types

type TimingStats struct {
	ElapsedTime int64 `json:"elapsed_time"`

	IterationTime *int64 `json:"iteration_time,omitempty"`
}

func (s *TimingStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTimingStats() *TimingStats { _ = "STUB: not implemented"; return nil }
