package types

type ShardsTotalStats struct {
	TotalCount int64 `json:"total_count"`
}

func (s *ShardsTotalStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShardsTotalStats() *ShardsTotalStats { _ = "STUB: not implemented"; return nil }
