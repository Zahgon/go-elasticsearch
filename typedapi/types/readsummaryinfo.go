package types

type ReadSummaryInfo struct {
	Count int `json:"count"`

	MaxWait Duration `json:"max_wait"`

	MaxWaitNanos int64 `json:"max_wait_nanos"`

	TotalElapsed Duration `json:"total_elapsed"`

	TotalElapsedNanos int64 `json:"total_elapsed_nanos"`

	TotalSize ByteSize `json:"total_size"`

	TotalSizeBytes int64 `json:"total_size_bytes"`

	TotalThrottled Duration `json:"total_throttled"`

	TotalThrottledNanos int64 `json:"total_throttled_nanos"`

	TotalWait Duration `json:"total_wait"`

	TotalWaitNanos int64 `json:"total_wait_nanos"`
}

func (s *ReadSummaryInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewReadSummaryInfo() *ReadSummaryInfo { _ = "STUB: not implemented"; return nil }
