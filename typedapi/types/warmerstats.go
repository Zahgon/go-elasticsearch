package types

type WarmerStats struct {
	Current           int64    `json:"current"`
	Total             int64    `json:"total"`
	TotalTime         Duration `json:"total_time,omitempty"`
	TotalTimeInMillis int64    `json:"total_time_in_millis"`
}

func (s *WarmerStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWarmerStats() *WarmerStats { _ = "STUB: not implemented"; return nil }
