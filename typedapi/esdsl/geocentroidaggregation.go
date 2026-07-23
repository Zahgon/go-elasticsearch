package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoCentroidAggregation struct {
	v *types.GeoCentroidAggregation
}

func NewGeoCentroidAggregation() *_geoCentroidAggregation { _ = "STUB: not implemented"; return nil }

func (s *_geoCentroidAggregation) Count(count int64) *_geoCentroidAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoCentroidAggregation) Location(geolocation types.GeoLocationVariant) *_geoCentroidAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoCentroidAggregation) Field(field string) *_geoCentroidAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoCentroidAggregation) Missing(missing types.MissingVariant) *_geoCentroidAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoCentroidAggregation) Script(script types.ScriptVariant) *_geoCentroidAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoCentroidAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoCentroidAggregation) GeoCentroidAggregationCaster() *types.GeoCentroidAggregation {
	_ = "STUB: not implemented"
	return nil
}
