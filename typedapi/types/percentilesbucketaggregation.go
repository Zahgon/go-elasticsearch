package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type PercentilesBucketAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`

	Percents []Float64 `json:"percents,omitempty"`
}

func (s *PercentilesBucketAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPercentilesBucketAggregation() *PercentilesBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

type PercentilesBucketAggregationVariant interface {
	PercentilesBucketAggregationCaster() *PercentilesBucketAggregation
}

func (s *PercentilesBucketAggregation) PercentilesBucketAggregationCaster() *PercentilesBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
