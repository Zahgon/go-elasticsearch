package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _latLonGeoLocation struct {
	v *types.LatLonGeoLocation
}

func NewLatLonGeoLocation(lat types.Float64, lon types.Float64) *_latLonGeoLocation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_latLonGeoLocation) Lat(lat types.Float64) *_latLonGeoLocation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_latLonGeoLocation) Lon(lon types.Float64) *_latLonGeoLocation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_latLonGeoLocation) LatLonGeoLocationCaster() *types.LatLonGeoLocation {
	_ = "STUB: not implemented"
	return nil
}
