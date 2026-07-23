package types

type NodeInfoTransport struct {
	BoundAddress   []string          `json:"bound_address"`
	Profiles       map[string]string `json:"profiles"`
	PublishAddress string            `json:"publish_address"`
}

func (s *NodeInfoTransport) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeInfoTransport() *NodeInfoTransport { _ = "STUB: not implemented"; return nil }
