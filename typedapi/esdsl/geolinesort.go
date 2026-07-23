package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoLineSort struct {
	v *types.GeoLineSort
}

func NewGeoLineSort() *_geoLineSort { _ = "STUB: not implemented"; return nil }

func (s *_geoLineSort) Field(field string) *_geoLineSort { _ = "STUB: not implemented"; return nil }

func (s *_geoLineSort) GeoLineSortCaster() *types.GeoLineSort {
	_ = "STUB: not implemented"
	return nil
}
