package types

type NodeInfoSettingsHttpType struct {
	Default string `json:"default"`
}

func (s *NodeInfoSettingsHttpType) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoSettingsHttpType() *NodeInfoSettingsHttpType { _ = "STUB: not implemented"; return nil }
