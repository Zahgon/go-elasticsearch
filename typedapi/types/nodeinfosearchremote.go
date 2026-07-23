package types

type NodeInfoSearchRemote struct {
	Connect string `json:"connect"`
}

func (s *NodeInfoSearchRemote) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoSearchRemote() *NodeInfoSearchRemote { _ = "STUB: not implemented"; return nil }
