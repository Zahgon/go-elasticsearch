package types

type IndicesIndexingPressureMemory struct {
	Limit *int `json:"limit,omitempty"`
}

func (s *IndicesIndexingPressureMemory) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndicesIndexingPressureMemory() *IndicesIndexingPressureMemory {
	_ = "STUB: not implemented"
	return nil
}

type IndicesIndexingPressureMemoryVariant interface {
	IndicesIndexingPressureMemoryCaster() *IndicesIndexingPressureMemory
}

func (s *IndicesIndexingPressureMemory) IndicesIndexingPressureMemoryCaster() *IndicesIndexingPressureMemory {
	_ = "STUB: not implemented"
	return nil
}
