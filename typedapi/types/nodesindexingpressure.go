package types

type NodesIndexingPressure struct {
	Memory *NodesIndexingPressureMemory `json:"memory,omitempty"`
}

func NewNodesIndexingPressure() *NodesIndexingPressure { _ = "STUB: not implemented"; return nil }
