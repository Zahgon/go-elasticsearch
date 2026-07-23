package types

type TransformNodeStats struct {
	Scheduler TransformSchedulerStats `json:"scheduler"`
}

func NewTransformNodeStats() *TransformNodeStats { _ = "STUB: not implemented"; return nil }
