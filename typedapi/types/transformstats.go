package types

type TransformStats struct {
	Checkpointing Checkpointing         `json:"checkpointing"`
	Health        *TransformStatsHealth `json:"health,omitempty"`
	Id            string                `json:"id"`
	Node          *NodeAttributes       `json:"node,omitempty"`
	Reason        *string               `json:"reason,omitempty"`
	State         string                `json:"state"`
	Stats         TransformIndexerStats `json:"stats"`
}

func (s *TransformStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTransformStats() *TransformStats { _ = "STUB: not implemented"; return nil }
