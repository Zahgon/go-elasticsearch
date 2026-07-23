package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type _geoLineAggregation struct {
	v *types.GeoLineAggregation
}

func NewGeoLineAggregation(point types.GeoLinePointVariant) *_geoLineAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoLineAggregation) IncludeSort(includesort bool) *_geoLineAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoLineAggregation) Point(point types.GeoLinePointVariant) *_geoLineAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoLineAggregation) Size(size int) *_geoLineAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoLineAggregation) Sort(sort types.GeoLineSortVariant) *_geoLineAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoLineAggregation) SortOrder(sortorder sortorder.SortOrder) *_geoLineAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoLineAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoLineAggregation) GeoLineAggregationCaster() *types.GeoLineAggregation {
	_ = "STUB: not implemented"
	return nil
}
