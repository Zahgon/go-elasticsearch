package types

type NodeInfoClient struct {
	Type string `json:"type"`
}

func (s *NodeInfoClient) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeInfoClient() *NodeInfoClient { _ = "STUB: not implemented"; return nil }
