package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _statsAggregation struct {
	v *types.StatsAggregation
}

func NewStatsAggregation() *_statsAggregation { _ = "STUB: not implemented"; return nil }

func (s *_statsAggregation) Field(field string) *_statsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_statsAggregation) Format(format string) *_statsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_statsAggregation) Missing(missing types.MissingVariant) *_statsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_statsAggregation) Script(script types.ScriptVariant) *_statsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_statsAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_statsAggregation) StatsAggregationCaster() *types.StatsAggregation {
	_ = "STUB: not implemented"
	return nil
}
