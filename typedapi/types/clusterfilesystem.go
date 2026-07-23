package types

type ClusterFileSystem struct {
	Available ByteSize `json:"available,omitempty"`

	AvailableInBytes           *int64   `json:"available_in_bytes,omitempty"`
	FloodStageFreeSpace        ByteSize `json:"flood_stage_free_space,omitempty"`
	FloodStageFreeSpaceInBytes *int64   `json:"flood_stage_free_space_in_bytes,omitempty"`

	Free ByteSize `json:"free,omitempty"`

	FreeInBytes                      *int64   `json:"free_in_bytes,omitempty"`
	FrozenFloodStageFreeSpace        ByteSize `json:"frozen_flood_stage_free_space,omitempty"`
	FrozenFloodStageFreeSpaceInBytes *int64   `json:"frozen_flood_stage_free_space_in_bytes,omitempty"`
	HighWatermarkFreeSpace           ByteSize `json:"high_watermark_free_space,omitempty"`
	HighWatermarkFreeSpaceInBytes    *int64   `json:"high_watermark_free_space_in_bytes,omitempty"`
	LowWatermarkFreeSpace            ByteSize `json:"low_watermark_free_space,omitempty"`
	LowWatermarkFreeSpaceInBytes     *int64   `json:"low_watermark_free_space_in_bytes,omitempty"`
	Mount                            *string  `json:"mount,omitempty"`
	Path                             *string  `json:"path,omitempty"`

	Total ByteSize `json:"total,omitempty"`

	TotalInBytes *int64  `json:"total_in_bytes,omitempty"`
	Type         *string `json:"type,omitempty"`
}

func (s *ClusterFileSystem) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClusterFileSystem() *ClusterFileSystem { _ = "STUB: not implemented"; return nil }
