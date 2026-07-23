package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoLinePoint struct {
	v *types.GeoLinePoint
}

func NewGeoLinePoint() *_geoLinePoint { _ = "STUB: not implemented"; return nil }

func (s *_geoLinePoint) Field(field string) *_geoLinePoint { _ = "STUB: not implemented"; return nil }

func (s *_geoLinePoint) GeoLinePointCaster() *types.GeoLinePoint {
	_ = "STUB: not implemented"
	return nil
}
