package types

type AutoFollowedCluster struct {
	ClusterName              string `json:"cluster_name"`
	LastSeenMetadataVersion  int64  `json:"last_seen_metadata_version"`
	TimeSinceLastCheckMillis int64  `json:"time_since_last_check_millis"`
}

func (s *AutoFollowedCluster) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAutoFollowedCluster() *AutoFollowedCluster { _ = "STUB: not implemented"; return nil }
