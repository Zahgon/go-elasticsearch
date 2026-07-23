package types

type NodeInfoXpackMl struct {
	UseAutoMachineMemoryPercent *bool `json:"use_auto_machine_memory_percent,omitempty"`
}

func (s *NodeInfoXpackMl) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeInfoXpackMl() *NodeInfoXpackMl { _ = "STUB: not implemented"; return nil }
