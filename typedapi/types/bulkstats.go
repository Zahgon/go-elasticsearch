package types

type BulkStats struct {
	AvgSize           ByteSize `json:"avg_size,omitempty"`
	AvgSizeInBytes    int64    `json:"avg_size_in_bytes"`
	AvgTime           Duration `json:"avg_time,omitempty"`
	AvgTimeInMillis   int64    `json:"avg_time_in_millis"`
	TotalOperations   int64    `json:"total_operations"`
	TotalSize         ByteSize `json:"total_size,omitempty"`
	TotalSizeInBytes  int64    `json:"total_size_in_bytes"`
	TotalTime         Duration `json:"total_time,omitempty"`
	TotalTimeInMillis int64    `json:"total_time_in_millis"`
}

func (s *BulkStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewBulkStats() *BulkStats { _ = "STUB: not implemented"; return nil }
