package types

type FollowIndexStats struct {
	Index string `json:"index"`

	Shards []CcrShardStats `json:"shards"`
}

func (s *FollowIndexStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFollowIndexStats() *FollowIndexStats { _ = "STUB: not implemented"; return nil }
