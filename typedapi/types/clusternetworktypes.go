package types

type ClusterNetworkTypes struct {
	HttpTypes map[string]int `json:"http_types"`

	TransportTypes map[string]int `json:"transport_types"`
}

func NewClusterNetworkTypes() *ClusterNetworkTypes { _ = "STUB: not implemented"; return nil }
