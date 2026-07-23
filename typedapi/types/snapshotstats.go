package types

type SnapshotStats struct {
	Incremental FileCountSnapshotStats `json:"incremental"`

	StartTimeInMillis int64    `json:"start_time_in_millis"`
	Time              Duration `json:"time,omitempty"`

	TimeInMillis int64 `json:"time_in_millis"`

	Total FileCountSnapshotStats `json:"total"`
}

func (s *SnapshotStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSnapshotStats() *SnapshotStats { _ = "STUB: not implemented"; return nil }
