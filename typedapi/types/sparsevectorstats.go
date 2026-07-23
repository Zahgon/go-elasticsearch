package types

type SparseVectorStats struct {
	ValueCount int64 `json:"value_count"`
}

func (s *SparseVectorStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSparseVectorStats() *SparseVectorStats { _ = "STUB: not implemented"; return nil }
