package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoHashLocation struct {
	v *types.GeoHashLocation
}

func NewGeoHashLocation() *_geoHashLocation { _ = "STUB: not implemented"; return nil }

func (s *_geoHashLocation) Geohash(geohash string) *_geoHashLocation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoHashLocation) GeoHashLocationCaster() *types.GeoHashLocation {
	_ = "STUB: not implemented"
	return nil
}
