package types

type Shared struct {
	BytesReadInBytes    ByteSize `json:"bytes_read_in_bytes"`
	BytesWrittenInBytes ByteSize `json:"bytes_written_in_bytes"`
	Evictions           int64    `json:"evictions"`
	NumRegions          int      `json:"num_regions"`
	Reads               int64    `json:"reads"`
	RegionSizeInBytes   ByteSize `json:"region_size_in_bytes"`
	SizeInBytes         ByteSize `json:"size_in_bytes"`
	Writes              int64    `json:"writes"`
}

func (s *Shared) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShared() *Shared { _ = "STUB: not implemented"; return nil }
