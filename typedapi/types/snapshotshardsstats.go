package types

type SnapshotShardsStats struct {
	Done int64 `json:"done"`

	Failed int64 `json:"failed"`

	Finalizing int64 `json:"finalizing"`

	Initializing int64 `json:"initializing"`

	Started int64 `json:"started"`

	Total int64 `json:"total"`
}

func (s *SnapshotShardsStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSnapshotShardsStats() *SnapshotShardsStats { _ = "STUB: not implemented"; return nil }
