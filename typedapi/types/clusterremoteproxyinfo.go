package types

type ClusterRemoteProxyInfo struct {
	ClusterCredentials *string `json:"cluster_credentials,omitempty"`

	Connected bool `json:"connected"`

	InitialConnectTimeout Duration `json:"initial_connect_timeout"`

	MaxProxySocketConnections int `json:"max_proxy_socket_connections"`

	Mode string `json:"mode,omitempty"`

	NumProxySocketsConnected int `json:"num_proxy_sockets_connected"`

	ProxyAddress string `json:"proxy_address"`
	ServerName   string `json:"server_name"`

	SkipUnavailable bool `json:"skip_unavailable"`
}

func (s *ClusterRemoteProxyInfo) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ClusterRemoteProxyInfo) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewClusterRemoteProxyInfo() *ClusterRemoteProxyInfo { _ = "STUB: not implemented"; return nil }
