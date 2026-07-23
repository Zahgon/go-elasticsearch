package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type GeoLineAggregation struct {
	IncludeSort *bool `json:"include_sort,omitempty"`

	Point GeoLinePoint `json:"point"`

	Size *int `json:"size,omitempty"`

	Sort *GeoLineSort `json:"sort,omitempty"`

	SortOrder *sortorder.SortOrder `json:"sort_order,omitempty"`
}

func (s *GeoLineAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoLineAggregation() *GeoLineAggregation { _ = "STUB: not implemented"; return nil }

type GeoLineAggregationVariant interface {
	GeoLineAggregationCaster() *GeoLineAggregation
}

func (s *GeoLineAggregation) GeoLineAggregationCaster() *GeoLineAggregation {
	_ = "STUB: not implemented"
	return nil
}
