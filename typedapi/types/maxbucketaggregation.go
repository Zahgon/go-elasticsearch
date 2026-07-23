package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type MaxBucketAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
}

func (s *MaxBucketAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMaxBucketAggregation() *MaxBucketAggregation { _ = "STUB: not implemented"; return nil }

type MaxBucketAggregationVariant interface {
	MaxBucketAggregationCaster() *MaxBucketAggregation
}

func (s *MaxBucketAggregation) MaxBucketAggregationCaster() *MaxBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
