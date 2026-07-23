package types

type OperatingSystemMemoryInfo struct {
	AdjustedTotal ByteSize `json:"adjusted_total,omitempty"`

	AdjustedTotalInBytes *int64 `json:"adjusted_total_in_bytes,omitempty"`

	Free ByteSize `json:"free,omitempty"`

	FreeInBytes int64 `json:"free_in_bytes"`

	FreePercent int `json:"free_percent"`

	Total ByteSize `json:"total,omitempty"`

	TotalInBytes int64 `json:"total_in_bytes"`

	Used ByteSize `json:"used,omitempty"`

	UsedInBytes int64 `json:"used_in_bytes"`

	UsedPercent int `json:"used_percent"`
}

func (s *OperatingSystemMemoryInfo) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewOperatingSystemMemoryInfo() *OperatingSystemMemoryInfo {
	_ = "STUB: not implemented"
	return nil
}
