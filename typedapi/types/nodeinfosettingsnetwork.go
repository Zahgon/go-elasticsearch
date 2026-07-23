package types

type NodeInfoSettingsNetwork struct {
	Host []string `json:"host,omitempty"`
}

func (s *NodeInfoSettingsNetwork) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoSettingsNetwork() *NodeInfoSettingsNetwork { _ = "STUB: not implemented"; return nil }
