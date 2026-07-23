package types

type NodeInfoBootstrap struct {
	MemoryLock string `json:"memory_lock"`
}

func (s *NodeInfoBootstrap) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeInfoBootstrap() *NodeInfoBootstrap { _ = "STUB: not implemented"; return nil }
