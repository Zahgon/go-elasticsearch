package types

type Allocations struct {
	CurrentDiskUsageInBytes *int64 `json:"current_disk_usage_in_bytes,omitempty"`

	ForecastedDiskUsageInBytes *int64 `json:"forecasted_disk_usage_in_bytes,omitempty"`

	ForecastedIngestLoad *Float64 `json:"forecasted_ingest_load,omitempty"`

	Shards *int `json:"shards,omitempty"`

	UndesiredShards *int `json:"undesired_shards,omitempty"`
}

func (s *Allocations) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAllocations() *Allocations { _ = "STUB: not implemented"; return nil }
