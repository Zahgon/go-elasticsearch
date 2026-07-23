package types

type CheckpointStats struct {
	Checkpoint           int64              `json:"checkpoint"`
	CheckpointProgress   *TransformProgress `json:"checkpoint_progress,omitempty"`
	TimeUpperBound       DateTime           `json:"time_upper_bound,omitempty"`
	TimeUpperBoundMillis *int64             `json:"time_upper_bound_millis,omitempty"`
	Timestamp            DateTime           `json:"timestamp,omitempty"`
	TimestampMillis      *int64             `json:"timestamp_millis,omitempty"`
}

func (s *CheckpointStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCheckpointStats() *CheckpointStats { _ = "STUB: not implemented"; return nil }
