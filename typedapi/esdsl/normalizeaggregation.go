package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/normalizemethod"
)

type _normalizeAggregation struct {
	v *types.NormalizeAggregation
}

func NewNormalizeAggregation() *_normalizeAggregation { _ = "STUB: not implemented"; return nil }

func (s *_normalizeAggregation) Method(method normalizemethod.NormalizeMethod) *_normalizeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_normalizeAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_normalizeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_normalizeAggregation) Format(format string) *_normalizeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_normalizeAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_normalizeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_normalizeAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_normalizeAggregation) NormalizeAggregationCaster() *types.NormalizeAggregation {
	_ = "STUB: not implemented"
	return nil
}
