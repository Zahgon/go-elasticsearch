package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _extendedStatsAggregation struct {
	v *types.ExtendedStatsAggregation
}

func NewExtendedStatsAggregation() *_extendedStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_extendedStatsAggregation) Sigma(sigma types.Float64) *_extendedStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_extendedStatsAggregation) Field(field string) *_extendedStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_extendedStatsAggregation) Format(format string) *_extendedStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_extendedStatsAggregation) Missing(missing types.MissingVariant) *_extendedStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_extendedStatsAggregation) Script(script types.ScriptVariant) *_extendedStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_extendedStatsAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_extendedStatsAggregation) ExtendedStatsAggregationCaster() *types.ExtendedStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}
