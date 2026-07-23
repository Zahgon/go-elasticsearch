package types

type SnapshotIndexStats struct {
	Shards      map[string]SnapshotShardsStatus `json:"shards"`
	ShardsStats SnapshotShardsStats             `json:"shards_stats"`
	Stats       SnapshotStats                   `json:"stats"`
}

func NewSnapshotIndexStats() *SnapshotIndexStats { _ = "STUB: not implemented"; return nil }
