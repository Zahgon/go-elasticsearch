package types

type NodeInfoSearch struct {
	Remote NodeInfoSearchRemote `json:"remote"`
}

func NewNodeInfoSearch() *NodeInfoSearch { _ = "STUB: not implemented"; return nil }
