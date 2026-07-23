package types

type IndicesIndexingPressure struct {
	Memory IndicesIndexingPressureMemory `json:"memory"`
}

func NewIndicesIndexingPressure() *IndicesIndexingPressure { _ = "STUB: not implemented"; return nil }

type IndicesIndexingPressureVariant interface {
	IndicesIndexingPressureCaster() *IndicesIndexingPressure
}

func (s *IndicesIndexingPressure) IndicesIndexingPressureCaster() *IndicesIndexingPressure {
	_ = "STUB: not implemented"
	return nil
}
