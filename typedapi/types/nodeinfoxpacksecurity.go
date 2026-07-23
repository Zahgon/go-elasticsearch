package types

type NodeInfoXpackSecurity struct {
	Authc     *NodeInfoXpackSecurityAuthc `json:"authc,omitempty"`
	Enabled   string                      `json:"enabled"`
	Http      *NodeInfoXpackSecuritySsl   `json:"http,omitempty"`
	Transport *NodeInfoXpackSecuritySsl   `json:"transport,omitempty"`
}

func (s *NodeInfoXpackSecurity) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoXpackSecurity() *NodeInfoXpackSecurity { _ = "STUB: not implemented"; return nil }
