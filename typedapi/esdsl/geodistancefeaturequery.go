package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoDistanceFeatureQuery struct {
	v *types.GeoDistanceFeatureQuery
}

func NewGeoDistanceFeatureQuery() *_geoDistanceFeatureQuery { _ = "STUB: not implemented"; return nil }

func (s *_geoDistanceFeatureQuery) Boost(boost float32) *_geoDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceFeatureQuery) Field(field string) *_geoDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceFeatureQuery) Origin(geolocation types.GeoLocationVariant) *_geoDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceFeatureQuery) Pivot(distance string) *_geoDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceFeatureQuery) QueryName_(queryname_ string) *_geoDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceFeatureQuery) QueryCaster() *types.Query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceFeatureQuery) GeoDistanceFeatureQueryCaster() *types.GeoDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}
