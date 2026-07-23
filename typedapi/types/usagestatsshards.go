package types

type UsageStatsShards struct {
	Routing                 ShardRouting       `json:"routing"`
	Stats                   IndicesShardsStats `json:"stats"`
	TrackingId              string             `json:"tracking_id"`
	TrackingStartedAtMillis int64              `json:"tracking_started_at_millis"`
}

func (s *UsageStatsShards) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewUsageStatsShards() *UsageStatsShards { _ = "STUB: not implemented"; return nil }
