package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _topMetricsAggregation struct {
	v *types.TopMetricsAggregation
}

func NewTopMetricsAggregation() *_topMetricsAggregation { _ = "STUB: not implemented"; return nil }

func (s *_topMetricsAggregation) Metrics(metrics ...types.TopMetricsValueVariant) *_topMetricsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topMetricsAggregation) Size(size int) *_topMetricsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topMetricsAggregation) Sort(sorts ...types.SortCombinationsVariant) *_topMetricsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topMetricsAggregation) SortValues(sortvalues []types.SortCombinations) *_topMetricsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topMetricsAggregation) Field(field string) *_topMetricsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topMetricsAggregation) Missing(missing types.MissingVariant) *_topMetricsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topMetricsAggregation) Script(script types.ScriptVariant) *_topMetricsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topMetricsAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topMetricsAggregation) TopMetricsAggregationCaster() *types.TopMetricsAggregation {
	_ = "STUB: not implemented"
	return nil
}
