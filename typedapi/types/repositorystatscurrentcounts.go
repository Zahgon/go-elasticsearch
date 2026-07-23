package types

type RepositoryStatsCurrentCounts struct {
	ActiveDeletions   int                   `json:"active_deletions"`
	Clones            int                   `json:"clones"`
	Deletions         int                   `json:"deletions"`
	Finalizations     int                   `json:"finalizations"`
	Shards            RepositoryStatsShards `json:"shards"`
	SnapshotDeletions int                   `json:"snapshot_deletions"`
	Snapshots         int                   `json:"snapshots"`
}

func (s *RepositoryStatsCurrentCounts) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRepositoryStatsCurrentCounts() *RepositoryStatsCurrentCounts {
	_ = "STUB: not implemented"
	return nil
}
