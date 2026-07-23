package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type CumulativeCardinalityAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
}

func (s *CumulativeCardinalityAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCumulativeCardinalityAggregation() *CumulativeCardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}

type CumulativeCardinalityAggregationVariant interface {
	CumulativeCardinalityAggregationCaster() *CumulativeCardinalityAggregation
}

func (s *CumulativeCardinalityAggregation) CumulativeCardinalityAggregationCaster() *CumulativeCardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}
