package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoPolygonPoints struct {
	v *types.GeoPolygonPoints
}

func NewGeoPolygonPoints() *_geoPolygonPoints { _ = "STUB: not implemented"; return nil }

func (s *_geoPolygonPoints) Points(points ...types.GeoLocationVariant) *_geoPolygonPoints {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPolygonPoints) PointsValues(pointsvalues []types.GeoLocation) *_geoPolygonPoints {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPolygonPoints) GeoPolygonPointsCaster() *types.GeoPolygonPoints {
	_ = "STUB: not implemented"
	return nil
}
