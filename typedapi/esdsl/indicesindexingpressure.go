package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indicesIndexingPressure struct {
	v *types.IndicesIndexingPressure
}

func NewIndicesIndexingPressure(memory types.IndicesIndexingPressureMemoryVariant) *_indicesIndexingPressure {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesIndexingPressure) Memory(memory types.IndicesIndexingPressureMemoryVariant) *_indicesIndexingPressure {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesIndexingPressure) IndicesIndexingPressureCaster() *types.IndicesIndexingPressure {
	_ = "STUB: not implemented"
	return nil
}
