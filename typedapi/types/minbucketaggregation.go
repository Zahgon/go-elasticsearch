package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type MinBucketAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
}

func (s *MinBucketAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMinBucketAggregation() *MinBucketAggregation { _ = "STUB: not implemented"; return nil }

type MinBucketAggregationVariant interface {
	MinBucketAggregationCaster() *MinBucketAggregation
}

func (s *MinBucketAggregation) MinBucketAggregationCaster() *MinBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
