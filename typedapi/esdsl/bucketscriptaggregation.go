package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _bucketScriptAggregation struct {
	v *types.BucketScriptAggregation
}

func NewBucketScriptAggregation() *_bucketScriptAggregation { _ = "STUB: not implemented"; return nil }

func (s *_bucketScriptAggregation) Script(script types.ScriptVariant) *_bucketScriptAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketScriptAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_bucketScriptAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketScriptAggregation) Format(format string) *_bucketScriptAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketScriptAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_bucketScriptAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketScriptAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketScriptAggregation) BucketScriptAggregationCaster() *types.BucketScriptAggregation {
	_ = "STUB: not implemented"
	return nil
}
