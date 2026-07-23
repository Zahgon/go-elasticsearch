package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoHashPrecision struct {
	v types.GeoHashPrecision
}

func NewGeoHashPrecision() *_geoHashPrecision { _ = "STUB: not implemented"; return nil }

func (u *_geoHashPrecision) Int(int int) *_geoHashPrecision { _ = "STUB: not implemented"; return nil }

func (u *_geoHashPrecision) String(string string) *_geoHashPrecision {
	_ = "STUB: not implemented"
	return nil
}

func (u *_geoHashPrecision) GeoHashPrecisionCaster() *types.GeoHashPrecision {
	_ = "STUB: not implemented"
	return nil
}
