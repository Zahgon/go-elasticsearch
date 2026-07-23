package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type BucketSelectorAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`

	Script *Script `json:"script,omitempty"`
}

func (s *BucketSelectorAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewBucketSelectorAggregation() *BucketSelectorAggregation {
	_ = "STUB: not implemented"
	return nil
}

type BucketSelectorAggregationVariant interface {
	BucketSelectorAggregationCaster() *BucketSelectorAggregation
}

func (s *BucketSelectorAggregation) BucketSelectorAggregationCaster() *BucketSelectorAggregation {
	_ = "STUB: not implemented"
	return nil
}
