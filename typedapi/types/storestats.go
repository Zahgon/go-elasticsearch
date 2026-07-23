package types

type StoreStats struct {
	Reserved ByteSize `json:"reserved,omitempty"`

	ReservedInBytes int64 `json:"reserved_in_bytes"`

	Size ByteSize `json:"size,omitempty"`

	SizeInBytes int64 `json:"size_in_bytes"`

	TotalDataSetSize ByteSize `json:"total_data_set_size,omitempty"`

	TotalDataSetSizeInBytes *int64 `json:"total_data_set_size_in_bytes,omitempty"`
}

func (s *StoreStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewStoreStats() *StoreStats { _ = "STUB: not implemented"; return nil }
