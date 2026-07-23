package types

type WriteSummaryInfo struct {
	Count int `json:"count"`

	TotalElapsed Duration `json:"total_elapsed"`

	TotalElapsedNanos int64 `json:"total_elapsed_nanos"`

	TotalSize ByteSize `json:"total_size"`

	TotalSizeBytes int64 `json:"total_size_bytes"`

	TotalThrottled Duration `json:"total_throttled"`

	TotalThrottledNanos int64 `json:"total_throttled_nanos"`
}

func (s *WriteSummaryInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWriteSummaryInfo() *WriteSummaryInfo { _ = "STUB: not implemented"; return nil }
