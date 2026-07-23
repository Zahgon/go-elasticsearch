package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _stringStatsAggregation struct {
	v *types.StringStatsAggregation
}

func NewStringStatsAggregation() *_stringStatsAggregation { _ = "STUB: not implemented"; return nil }

func (s *_stringStatsAggregation) ShowDistribution(showdistribution bool) *_stringStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stringStatsAggregation) Field(field string) *_stringStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stringStatsAggregation) Missing(missing types.MissingVariant) *_stringStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stringStatsAggregation) Script(script types.ScriptVariant) *_stringStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stringStatsAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stringStatsAggregation) StringStatsAggregationCaster() *types.StringStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}
