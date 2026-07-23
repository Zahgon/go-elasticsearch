package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _cartesianBoundsAggregation struct {
	v *types.CartesianBoundsAggregation
}

func NewCartesianBoundsAggregation() *_cartesianBoundsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cartesianBoundsAggregation) Field(field string) *_cartesianBoundsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cartesianBoundsAggregation) Missing(missing types.MissingVariant) *_cartesianBoundsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cartesianBoundsAggregation) Script(script types.ScriptVariant) *_cartesianBoundsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cartesianBoundsAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cartesianBoundsAggregation) CartesianBoundsAggregationCaster() *types.CartesianBoundsAggregation {
	_ = "STUB: not implemented"
	return nil
}
