package types

type FlushStats struct {
	Periodic          int64    `json:"periodic"`
	Total             int64    `json:"total"`
	TotalTime         Duration `json:"total_time,omitempty"`
	TotalTimeInMillis int64    `json:"total_time_in_millis"`
}

func (s *FlushStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFlushStats() *FlushStats { _ = "STUB: not implemented"; return nil }
