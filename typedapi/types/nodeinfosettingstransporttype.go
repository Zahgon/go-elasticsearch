package types

type NodeInfoSettingsTransportType struct {
	Default string `json:"default"`
}

func (s *NodeInfoSettingsTransportType) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoSettingsTransportType() *NodeInfoSettingsTransportType {
	_ = "STUB: not implemented"
	return nil
}
