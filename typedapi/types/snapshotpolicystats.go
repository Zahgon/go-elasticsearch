package types

type SnapshotPolicyStats struct {
	Policy                   string `json:"policy"`
	SnapshotDeletionFailures int64  `json:"snapshot_deletion_failures"`
	SnapshotsDeleted         int64  `json:"snapshots_deleted"`
	SnapshotsFailed          int64  `json:"snapshots_failed"`
	SnapshotsTaken           int64  `json:"snapshots_taken"`
}

func (s *SnapshotPolicyStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSnapshotPolicyStats() *SnapshotPolicyStats { _ = "STUB: not implemented"; return nil }
