package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type _geoDistanceSort struct {
	v *types.GeoDistanceSort
}

func NewGeoDistanceSort() *_geoDistanceSort { _ = "STUB: not implemented"; return nil }

func (s *_geoDistanceSort) DistanceType(distancetype geodistancetype.GeoDistanceType) *_geoDistanceSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceSort) GeoDistanceSort(geodistancesort map[string][]types.GeoLocation) *_geoDistanceSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceSort) IgnoreUnmapped(ignoreunmapped bool) *_geoDistanceSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceSort) Mode(mode sortmode.SortMode) *_geoDistanceSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceSort) Nested(nested types.NestedSortValueVariant) *_geoDistanceSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceSort) Order(order sortorder.SortOrder) *_geoDistanceSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceSort) Unit(unit distanceunit.DistanceUnit) *_geoDistanceSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceSort) SortOptionsCaster() *types.SortOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceSort) GeoDistanceSortCaster() *types.GeoDistanceSort {
	_ = "STUB: not implemented"
	return nil
}
