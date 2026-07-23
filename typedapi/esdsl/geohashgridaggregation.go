package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoHashGridAggregation struct {
	v *types.GeoHashGridAggregation
}

func NewGeoHashGridAggregation() *_geoHashGridAggregation { _ = "STUB: not implemented"; return nil }

func (s *_geoHashGridAggregation) Bounds(geobounds types.GeoBoundsVariant) *_geoHashGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoHashGridAggregation) Field(field string) *_geoHashGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoHashGridAggregation) Precision(geohashprecision types.GeoHashPrecisionVariant) *_geoHashGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoHashGridAggregation) ShardSize(shardsize int) *_geoHashGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoHashGridAggregation) Size(size int) *_geoHashGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoHashGridAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoHashGridAggregation) GeoHashGridAggregationCaster() *types.GeoHashGridAggregation {
	_ = "STUB: not implemented"
	return nil
}
