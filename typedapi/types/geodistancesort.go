package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type GeoDistanceSort struct {
	DistanceType    *geodistancetype.GeoDistanceType `json:"distance_type,omitempty"`
	GeoDistanceSort map[string][]GeoLocation         `json:"-"`
	IgnoreUnmapped  *bool                            `json:"ignore_unmapped,omitempty"`
	Mode            *sortmode.SortMode               `json:"mode,omitempty"`
	Nested          *NestedSortValue                 `json:"nested,omitempty"`
	Order           *sortorder.SortOrder             `json:"order,omitempty"`
	Unit            *distanceunit.DistanceUnit       `json:"unit,omitempty"`
}

func (s *GeoDistanceSort) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GeoDistanceSort) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGeoDistanceSort() *GeoDistanceSort { _ = "STUB: not implemented"; return nil }

type GeoDistanceSortVariant interface {
	GeoDistanceSortCaster() *GeoDistanceSort
}

func (s *GeoDistanceSort) GeoDistanceSortCaster() *GeoDistanceSort {
	_ = "STUB: not implemented"
	return nil
}
