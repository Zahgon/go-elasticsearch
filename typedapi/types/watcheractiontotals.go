package types

type WatcherActionTotals struct {
	Total         Duration `json:"total"`
	TotalTimeInMs int64    `json:"total_time_in_ms"`
}

func (s *WatcherActionTotals) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewWatcherActionTotals() *WatcherActionTotals { _ = "STUB: not implemented"; return nil }
