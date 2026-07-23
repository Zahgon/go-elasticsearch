package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type StatsBucketAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
}

func (s *StatsBucketAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewStatsBucketAggregation() *StatsBucketAggregation { _ = "STUB: not implemented"; return nil }

type StatsBucketAggregationVariant interface {
	StatsBucketAggregationCaster() *StatsBucketAggregation
}

func (s *StatsBucketAggregation) StatsBucketAggregationCaster() *StatsBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
