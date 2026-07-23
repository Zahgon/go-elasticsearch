package types

type NodeInfoXpackSecurityAuthcToken struct {
	Enabled string `json:"enabled"`
}

func (s *NodeInfoXpackSecurityAuthcToken) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoXpackSecurityAuthcToken() *NodeInfoXpackSecurityAuthcToken {
	_ = "STUB: not implemented"
	return nil
}
