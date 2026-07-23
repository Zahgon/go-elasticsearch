package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _filtersAggregation struct {
	v *types.FiltersAggregation
}

func NewFiltersAggregation() *_filtersAggregation { _ = "STUB: not implemented"; return nil }

func (s *_filtersAggregation) Filters(bucketsquery types.BucketsQueryVariant) *_filtersAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filtersAggregation) Keyed(keyed bool) *_filtersAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filtersAggregation) OtherBucket(otherbucket bool) *_filtersAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filtersAggregation) OtherBucketKey(otherbucketkey string) *_filtersAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filtersAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filtersAggregation) FiltersAggregationCaster() *types.FiltersAggregation {
	_ = "STUB: not implemented"
	return nil
}
