package types

type NodeInfoXpackSecurityAuthcRealms struct {
	File   map[string]NodeInfoXpackSecurityAuthcRealmsStatus `json:"file,omitempty"`
	Native map[string]NodeInfoXpackSecurityAuthcRealmsStatus `json:"native,omitempty"`
	Pki    map[string]NodeInfoXpackSecurityAuthcRealmsStatus `json:"pki,omitempty"`
}

func NewNodeInfoXpackSecurityAuthcRealms() *NodeInfoXpackSecurityAuthcRealms {
	_ = "STUB: not implemented"
	return nil
}
