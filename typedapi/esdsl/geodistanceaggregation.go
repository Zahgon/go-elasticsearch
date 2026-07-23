package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geodistancetype"
)

type _geoDistanceAggregation struct {
	v *types.GeoDistanceAggregation
}

func NewGeoDistanceAggregation() *_geoDistanceAggregation { _ = "STUB: not implemented"; return nil }

func (s *_geoDistanceAggregation) DistanceType(distancetype geodistancetype.GeoDistanceType) *_geoDistanceAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceAggregation) Field(field string) *_geoDistanceAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceAggregation) Origin(geolocation types.GeoLocationVariant) *_geoDistanceAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceAggregation) Ranges(ranges ...types.AggregationRangeVariant) *_geoDistanceAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceAggregation) RangesValues(rangesvalues []types.AggregationRange) *_geoDistanceAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceAggregation) Unit(unit distanceunit.DistanceUnit) *_geoDistanceAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceAggregation) GeoDistanceAggregationCaster() *types.GeoDistanceAggregation {
	_ = "STUB: not implemented"
	return nil
}
