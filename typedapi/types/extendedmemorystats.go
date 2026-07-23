package types

type ExtendedMemoryStats struct {
	AdjustedTotalInBytes *int64 `json:"adjusted_total_in_bytes,omitempty"`

	FreeInBytes *int64 `json:"free_in_bytes,omitempty"`

	FreePercent     *int    `json:"free_percent,omitempty"`
	Resident        *string `json:"resident,omitempty"`
	ResidentInBytes *int64  `json:"resident_in_bytes,omitempty"`
	Share           *string `json:"share,omitempty"`
	ShareInBytes    *int64  `json:"share_in_bytes,omitempty"`

	TotalInBytes        *int64  `json:"total_in_bytes,omitempty"`
	TotalVirtual        *string `json:"total_virtual,omitempty"`
	TotalVirtualInBytes *int64  `json:"total_virtual_in_bytes,omitempty"`

	UsedInBytes *int64 `json:"used_in_bytes,omitempty"`

	UsedPercent *int `json:"used_percent,omitempty"`
}

func (s *ExtendedMemoryStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewExtendedMemoryStats() *ExtendedMemoryStats { _ = "STUB: not implemented"; return nil }
