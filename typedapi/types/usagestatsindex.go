package types

type UsageStatsIndex struct {
	Shards []UsageStatsShards `json:"shards"`
}

func NewUsageStatsIndex() *UsageStatsIndex { _ = "STUB: not implemented"; return nil }
