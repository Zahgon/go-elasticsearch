package types

type DataPathStats struct {
	Available *string `json:"available,omitempty"`

	AvailableInBytes     *int64  `json:"available_in_bytes,omitempty"`
	DiskQueue            *string `json:"disk_queue,omitempty"`
	DiskReadSize         *string `json:"disk_read_size,omitempty"`
	DiskReadSizeInBytes  *int64  `json:"disk_read_size_in_bytes,omitempty"`
	DiskReads            *int64  `json:"disk_reads,omitempty"`
	DiskWriteSize        *string `json:"disk_write_size,omitempty"`
	DiskWriteSizeInBytes *int64  `json:"disk_write_size_in_bytes,omitempty"`
	DiskWrites           *int64  `json:"disk_writes,omitempty"`

	Free *string `json:"free,omitempty"`

	FreeInBytes *int64 `json:"free_in_bytes,omitempty"`

	Mount *string `json:"mount,omitempty"`

	Path *string `json:"path,omitempty"`

	Total *string `json:"total,omitempty"`

	TotalInBytes *int64 `json:"total_in_bytes,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (s *DataPathStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDataPathStats() *DataPathStats { _ = "STUB: not implemented"; return nil }
