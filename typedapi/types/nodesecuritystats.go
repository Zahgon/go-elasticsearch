package types

type NodeSecurityStats struct {
	Roles RolesStats `json:"roles"`
}

func NewNodeSecurityStats() *NodeSecurityStats { _ = "STUB: not implemented"; return nil }
