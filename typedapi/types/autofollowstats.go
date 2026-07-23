package types

type AutoFollowStats struct {
	AutoFollowedClusters []AutoFollowedCluster `json:"auto_followed_clusters"`

	NumberOfFailedFollowIndices int64 `json:"number_of_failed_follow_indices"`

	NumberOfFailedRemoteClusterStateRequests int64 `json:"number_of_failed_remote_cluster_state_requests"`

	NumberOfSuccessfulFollowIndices int64 `json:"number_of_successful_follow_indices"`

	RecentAutoFollowErrors []ErrorCause `json:"recent_auto_follow_errors"`
}

func (s *AutoFollowStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAutoFollowStats() *AutoFollowStats { _ = "STUB: not implemented"; return nil }
