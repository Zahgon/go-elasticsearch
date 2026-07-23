package types

type NodeInfoXpackSecurityAuthcRealmsStatus struct {
	Enabled *string `json:"enabled,omitempty"`
	Order   string  `json:"order"`
}

func (s *NodeInfoXpackSecurityAuthcRealmsStatus) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoXpackSecurityAuthcRealmsStatus() *NodeInfoXpackSecurityAuthcRealmsStatus {
	_ = "STUB: not implemented"
	return nil
}
