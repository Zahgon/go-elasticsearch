package types

type SnapshotShardFailure struct {
	Index     string  `json:"index"`
	IndexUuid string  `json:"index_uuid"`
	NodeId    *string `json:"node_id,omitempty"`
	Reason    string  `json:"reason"`
	ShardId   int     `json:"shard_id"`
	Status    string  `json:"status"`
}

func (s *SnapshotShardFailure) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSnapshotShardFailure() *SnapshotShardFailure { _ = "STUB: not implemented"; return nil }
