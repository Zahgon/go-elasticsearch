package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoLocation struct {
	v types.GeoLocation
}

func NewGeoLocation() *_geoLocation { _ = "STUB: not implemented"; return nil }

func (u *_geoLocation) LatLonGeoLocation(latlongeolocation types.LatLonGeoLocationVariant) *_geoLocation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_latLonGeoLocation) GeoLocationCaster() *types.GeoLocation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_geoLocation) GeoHashLocation(geohashlocation types.GeoHashLocationVariant) *_geoLocation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_geoHashLocation) GeoLocationCaster() *types.GeoLocation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_geoLocation) Doubles(doubles ...types.Float64) *_geoLocation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_geoLocation) String(string string) *_geoLocation { _ = "STUB: not implemented"; return nil }

func (u *_geoLocation) GeoLocationCaster() *types.GeoLocation {
	_ = "STUB: not implemented"
	return nil
}
