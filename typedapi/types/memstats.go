package types

type MemStats struct {
	AdjustedTotal ByteSize `json:"adjusted_total,omitempty"`

	AdjustedTotalInBytes int `json:"adjusted_total_in_bytes"`

	Ml MemMlStats `json:"ml"`

	Total ByteSize `json:"total,omitempty"`

	TotalInBytes int `json:"total_in_bytes"`
}

func (s *MemStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMemStats() *MemStats { _ = "STUB: not implemented"; return nil }
