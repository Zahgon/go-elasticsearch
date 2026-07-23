package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _bucketSortAggregation struct {
	v *types.BucketSortAggregation
}

func NewBucketSortAggregation() *_bucketSortAggregation { _ = "STUB: not implemented"; return nil }

func (s *_bucketSortAggregation) From(from int) *_bucketSortAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketSortAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_bucketSortAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketSortAggregation) Size(size int) *_bucketSortAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketSortAggregation) Sort(sorts ...types.SortCombinationsVariant) *_bucketSortAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketSortAggregation) SortValues(sortvalues []types.SortCombinations) *_bucketSortAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketSortAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketSortAggregation) BucketSortAggregationCaster() *types.BucketSortAggregation {
	_ = "STUB: not implemented"
	return nil
}
