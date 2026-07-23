package types

type NodesIndexingPressureMemory struct {
	Current *PressureMemory `json:"current,omitempty"`

	Limit ByteSize `json:"limit,omitempty"`

	LimitInBytes *int64 `json:"limit_in_bytes,omitempty"`

	Total *PressureMemory `json:"total,omitempty"`
}

func (s *NodesIndexingPressureMemory) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodesIndexingPressureMemory() *NodesIndexingPressureMemory {
	_ = "STUB: not implemented"
	return nil
}
