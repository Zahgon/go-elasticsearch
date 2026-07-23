package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortmode"
)

type _matrixStatsAggregation struct {
	v *types.MatrixStatsAggregation
}

func NewMatrixStatsAggregation() *_matrixStatsAggregation { _ = "STUB: not implemented"; return nil }

func (s *_matrixStatsAggregation) Mode(mode sortmode.SortMode) *_matrixStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matrixStatsAggregation) Fields(fields ...string) *_matrixStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matrixStatsAggregation) Missing(missing map[string]types.Float64) *_matrixStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matrixStatsAggregation) AddMissing(key string, value types.Float64) *_matrixStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matrixStatsAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matrixStatsAggregation) MatrixStatsAggregationCaster() *types.MatrixStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}
