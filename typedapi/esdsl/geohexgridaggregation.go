package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geohexGridAggregation struct {
	v *types.GeohexGridAggregation
}

func NewGeohexGridAggregation() *_geohexGridAggregation { _ = "STUB: not implemented"; return nil }

func (s *_geohexGridAggregation) Bounds(geobounds types.GeoBoundsVariant) *_geohexGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geohexGridAggregation) Field(field string) *_geohexGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geohexGridAggregation) Precision(precision int) *_geohexGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geohexGridAggregation) ShardSize(shardsize int) *_geohexGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geohexGridAggregation) Size(size int) *_geohexGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geohexGridAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geohexGridAggregation) GeohexGridAggregationCaster() *types.GeohexGridAggregation {
	_ = "STUB: not implemented"
	return nil
}
