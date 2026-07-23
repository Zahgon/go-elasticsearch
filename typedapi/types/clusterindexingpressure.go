package types

type ClusterIndexingPressure struct {
	Memory NodesIndexingPressureMemory `json:"memory"`
}

func NewClusterIndexingPressure() *ClusterIndexingPressure { _ = "STUB: not implemented"; return nil }
