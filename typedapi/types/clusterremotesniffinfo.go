package types

type ClusterRemoteSniffInfo struct {
	Connected bool `json:"connected"`

	InitialConnectTimeout Duration `json:"initial_connect_timeout"`

	MaxConnectionsPerCluster int `json:"max_connections_per_cluster"`

	Mode string `json:"mode,omitempty"`

	NumNodesConnected int64 `json:"num_nodes_connected"`

	Seeds []string `json:"seeds"`

	SkipUnavailable bool `json:"skip_unavailable"`
}

func (s *ClusterRemoteSniffInfo) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ClusterRemoteSniffInfo) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewClusterRemoteSniffInfo() *ClusterRemoteSniffInfo { _ = "STUB: not implemented"; return nil }
