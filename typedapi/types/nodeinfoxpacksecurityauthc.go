package types

type NodeInfoXpackSecurityAuthc struct {
	Realms *NodeInfoXpackSecurityAuthcRealms `json:"realms,omitempty"`
	Token  *NodeInfoXpackSecurityAuthcToken  `json:"token,omitempty"`
}

func NewNodeInfoXpackSecurityAuthc() *NodeInfoXpackSecurityAuthc {
	_ = "STUB: not implemented"
	return nil
}
