package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type ExtendedStatsBucketAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`

	Sigma *Float64 `json:"sigma,omitempty"`
}

func (s *ExtendedStatsBucketAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewExtendedStatsBucketAggregation() *ExtendedStatsBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

type ExtendedStatsBucketAggregationVariant interface {
	ExtendedStatsBucketAggregationCaster() *ExtendedStatsBucketAggregation
}

func (s *ExtendedStatsBucketAggregation) ExtendedStatsBucketAggregationCaster() *ExtendedStatsBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
