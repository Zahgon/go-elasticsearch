package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type AverageBucketAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
}

func (s *AverageBucketAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAverageBucketAggregation() *AverageBucketAggregation { _ = "STUB: not implemented"; return nil }

type AverageBucketAggregationVariant interface {
	AverageBucketAggregationCaster() *AverageBucketAggregation
}

func (s *AverageBucketAggregation) AverageBucketAggregationCaster() *AverageBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
