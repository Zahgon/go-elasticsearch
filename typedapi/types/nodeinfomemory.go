package types

type NodeInfoMemory struct {
	Total        string `json:"total"`
	TotalInBytes int64  `json:"total_in_bytes"`
}

func (s *NodeInfoMemory) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeInfoMemory() *NodeInfoMemory { _ = "STUB: not implemented"; return nil }
