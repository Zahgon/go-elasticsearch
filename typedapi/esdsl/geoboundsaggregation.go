package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoBoundsAggregation struct {
	v *types.GeoBoundsAggregation
}

func NewGeoBoundsAggregation() *_geoBoundsAggregation { _ = "STUB: not implemented"; return nil }

func (s *_geoBoundsAggregation) WrapLongitude(wraplongitude bool) *_geoBoundsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoBoundsAggregation) Field(field string) *_geoBoundsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoBoundsAggregation) Missing(missing types.MissingVariant) *_geoBoundsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoBoundsAggregation) Script(script types.ScriptVariant) *_geoBoundsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoBoundsAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoBoundsAggregation) GeoBoundsAggregationCaster() *types.GeoBoundsAggregation {
	_ = "STUB: not implemented"
	return nil
}
