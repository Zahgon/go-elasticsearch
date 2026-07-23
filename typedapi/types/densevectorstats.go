package types

type DenseVectorStats struct {
	OffHeap    *DenseVectorOffHeapStats `json:"off_heap,omitempty"`
	ValueCount int64                    `json:"value_count"`
}

func (s *DenseVectorStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDenseVectorStats() *DenseVectorStats { _ = "STUB: not implemented"; return nil }
