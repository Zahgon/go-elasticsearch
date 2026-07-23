package types

type FollowStats struct {
	Indices []FollowIndexStats `json:"indices"`
}

func NewFollowStats() *FollowStats { _ = "STUB: not implemented"; return nil }
