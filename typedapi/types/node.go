package types

type Node struct {
	SharedCache Shared `json:"shared_cache"`
}

func NewNode() *Node { _ = "STUB: not implemented"; return nil }
