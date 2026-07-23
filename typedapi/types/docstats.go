package types

type DocStats struct {
	Count int64 `json:"count"`

	Deleted *int64 `json:"deleted,omitempty"`

	TotalSize ByteSize `json:"total_size,omitempty"`

	TotalSizeInBytes int64 `json:"total_size_in_bytes"`
}

func (s *DocStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDocStats() *DocStats { _ = "STUB: not implemented"; return nil }
