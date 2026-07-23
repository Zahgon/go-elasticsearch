package types

type ClusterSnapshotStats struct {
	CurrentCounts SnapshotCurrentCounts         `json:"current_counts"`
	Repositories  map[string]PerRepositoryStats `json:"repositories"`
}

func NewClusterSnapshotStats() *ClusterSnapshotStats { _ = "STUB: not implemented"; return nil }
