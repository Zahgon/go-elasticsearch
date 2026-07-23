package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoTileGridAggregation struct {
	v *types.GeoTileGridAggregation
}

func NewGeoTileGridAggregation() *_geoTileGridAggregation { _ = "STUB: not implemented"; return nil }

func (s *_geoTileGridAggregation) Bounds(geobounds types.GeoBoundsVariant) *_geoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoTileGridAggregation) Field(field string) *_geoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoTileGridAggregation) Precision(geotileprecision int) *_geoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoTileGridAggregation) ShardSize(shardsize int) *_geoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoTileGridAggregation) Size(size int) *_geoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoTileGridAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoTileGridAggregation) PivotGroupByContainerCaster() *types.PivotGroupByContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoTileGridAggregation) GeoTileGridAggregationCaster() *types.GeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}
