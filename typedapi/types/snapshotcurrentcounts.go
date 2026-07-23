package types

type SnapshotCurrentCounts struct {
	Cleanups int `json:"cleanups"`

	ConcurrentOperations int `json:"concurrent_operations"`

	ShardSnapshots int `json:"shard_snapshots"`

	SnapshotDeletions int `json:"snapshot_deletions"`

	Snapshots int `json:"snapshots"`
}

func (s *SnapshotCurrentCounts) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSnapshotCurrentCounts() *SnapshotCurrentCounts { _ = "STUB: not implemented"; return nil }
