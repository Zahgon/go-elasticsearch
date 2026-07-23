package types

type NodeThreadPoolInfo struct {
	Core      *int     `json:"core,omitempty"`
	KeepAlive Duration `json:"keep_alive,omitempty"`
	Max       *int     `json:"max,omitempty"`
	QueueSize int      `json:"queue_size"`
	Size      *int     `json:"size,omitempty"`
	Type      string   `json:"type"`
}

func (s *NodeThreadPoolInfo) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeThreadPoolInfo() *NodeThreadPoolInfo { _ = "STUB: not implemented"; return nil }
