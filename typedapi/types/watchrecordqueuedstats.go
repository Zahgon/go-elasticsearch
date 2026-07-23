package types

type WatchRecordQueuedStats struct {
	ExecutionTime DateTime `json:"execution_time"`
}

func (s *WatchRecordQueuedStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewWatchRecordQueuedStats() *WatchRecordQueuedStats { _ = "STUB: not implemented"; return nil }
