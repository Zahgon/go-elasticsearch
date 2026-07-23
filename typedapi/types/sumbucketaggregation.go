package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type SumBucketAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
}

func (s *SumBucketAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSumBucketAggregation() *SumBucketAggregation { _ = "STUB: not implemented"; return nil }

type SumBucketAggregationVariant interface {
	SumBucketAggregationCaster() *SumBucketAggregation
}

func (s *SumBucketAggregation) SumBucketAggregationCaster() *SumBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
