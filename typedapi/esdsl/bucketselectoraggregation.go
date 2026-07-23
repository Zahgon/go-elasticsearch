package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _bucketSelectorAggregation struct {
	v *types.BucketSelectorAggregation
}

func NewBucketSelectorAggregation() *_bucketSelectorAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketSelectorAggregation) Script(script types.ScriptVariant) *_bucketSelectorAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketSelectorAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_bucketSelectorAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketSelectorAggregation) Format(format string) *_bucketSelectorAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketSelectorAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_bucketSelectorAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketSelectorAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketSelectorAggregation) BucketSelectorAggregationCaster() *types.BucketSelectorAggregation {
	_ = "STUB: not implemented"
	return nil
}
