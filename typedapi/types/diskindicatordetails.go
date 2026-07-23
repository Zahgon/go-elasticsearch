package types

type DiskIndicatorDetails struct {
	IndicesWithReadonlyBlock     int64 `json:"indices_with_readonly_block"`
	NodesOverFloodStageWatermark int64 `json:"nodes_over_flood_stage_watermark"`
	NodesOverHighWatermark       int64 `json:"nodes_over_high_watermark"`
	NodesWithEnoughDiskSpace     int64 `json:"nodes_with_enough_disk_space"`
	NodesWithUnknownDiskStatus   int64 `json:"nodes_with_unknown_disk_status"`
}

func (s *DiskIndicatorDetails) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDiskIndicatorDetails() *DiskIndicatorDetails { _ = "STUB: not implemented"; return nil }
