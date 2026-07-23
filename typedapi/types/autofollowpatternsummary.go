package types

type AutoFollowPatternSummary struct {
	Active bool `json:"active"`

	FollowIndexPattern *string `json:"follow_index_pattern,omitempty"`

	LeaderIndexExclusionPatterns []string `json:"leader_index_exclusion_patterns"`

	LeaderIndexPatterns []string `json:"leader_index_patterns"`

	MaxOutstandingReadRequests int `json:"max_outstanding_read_requests"`

	RemoteCluster string `json:"remote_cluster"`
}

func (s *AutoFollowPatternSummary) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAutoFollowPatternSummary() *AutoFollowPatternSummary { _ = "STUB: not implemented"; return nil }
