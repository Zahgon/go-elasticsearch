package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geodistancetype"
)

type GeoDistanceAggregation struct {
	DistanceType *geodistancetype.GeoDistanceType `json:"distance_type,omitempty"`

	Field *string `json:"field,omitempty"`

	Origin GeoLocation `json:"origin,omitempty"`

	Ranges []AggregationRange `json:"ranges,omitempty"`

	Unit *distanceunit.DistanceUnit `json:"unit,omitempty"`
}

func (s *GeoDistanceAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoDistanceAggregation() *GeoDistanceAggregation { _ = "STUB: not implemented"; return nil }

type GeoDistanceAggregationVariant interface {
	GeoDistanceAggregationCaster() *GeoDistanceAggregation
}

func (s *GeoDistanceAggregation) GeoDistanceAggregationCaster() *GeoDistanceAggregation {
	_ = "STUB: not implemented"
	return nil
}
