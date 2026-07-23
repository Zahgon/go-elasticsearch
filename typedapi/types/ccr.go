package types

type Ccr struct {
	AutoFollowPatternsCount int  `json:"auto_follow_patterns_count"`
	Available               bool `json:"available"`
	Enabled                 bool `json:"enabled"`
	FollowerIndicesCount    int  `json:"follower_indices_count"`
}

func (s *Ccr) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCcr() *Ccr { _ = "STUB: not implemented"; return nil }
