package types

type RefreshStats struct {
	ExternalTotal             int64    `json:"external_total"`
	ExternalTotalTimeInMillis int64    `json:"external_total_time_in_millis"`
	Listeners                 int64    `json:"listeners"`
	Total                     int64    `json:"total"`
	TotalTime                 Duration `json:"total_time,omitempty"`
	TotalTimeInMillis         int64    `json:"total_time_in_millis"`
}

func (s *RefreshStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRefreshStats() *RefreshStats { _ = "STUB: not implemented"; return nil }
