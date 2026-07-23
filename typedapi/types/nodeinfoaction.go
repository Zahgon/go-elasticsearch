package types

type NodeInfoAction struct {
	DestructiveRequiresName string `json:"destructive_requires_name"`
}

func (s *NodeInfoAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeInfoAction() *NodeInfoAction { _ = "STUB: not implemented"; return nil }
