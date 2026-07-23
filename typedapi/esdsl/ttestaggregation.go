package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ttesttype"
)

type _tTestAggregation struct {
	v *types.TTestAggregation
}

func NewTTestAggregation() *_tTestAggregation { _ = "STUB: not implemented"; return nil }

func (s *_tTestAggregation) A(a types.TestPopulationVariant) *_tTestAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tTestAggregation) B(b types.TestPopulationVariant) *_tTestAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tTestAggregation) Type(type_ ttesttype.TTestType) *_tTestAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tTestAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tTestAggregation) TTestAggregationCaster() *types.TTestAggregation {
	_ = "STUB: not implemented"
	return nil
}
