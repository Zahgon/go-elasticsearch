package types

type ShardRetentionLeases struct {
	Leases      []ShardLease `json:"leases"`
	PrimaryTerm int64        `json:"primary_term"`
	Version     int64        `json:"version"`
}

func (s *ShardRetentionLeases) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewShardRetentionLeases() *ShardRetentionLeases { _ = "STUB: not implemented"; return nil }
