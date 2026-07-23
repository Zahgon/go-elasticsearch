package types

type TransformSchedulerStats struct {
	PeekTransform            *string `json:"peek_transform,omitempty"`
	RegisteredTransformCount int     `json:"registered_transform_count"`
}

func (s *TransformSchedulerStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTransformSchedulerStats() *TransformSchedulerStats { _ = "STUB: not implemented"; return nil }
