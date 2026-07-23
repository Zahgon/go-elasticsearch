package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type SerialDifferencingAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`

	Lag *int `json:"lag,omitempty"`
}

func (s *SerialDifferencingAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSerialDifferencingAggregation() *SerialDifferencingAggregation {
	_ = "STUB: not implemented"
	return nil
}

type SerialDifferencingAggregationVariant interface {
	SerialDifferencingAggregationCaster() *SerialDifferencingAggregation
}

func (s *SerialDifferencingAggregation) SerialDifferencingAggregationCaster() *SerialDifferencingAggregation {
	_ = "STUB: not implemented"
	return nil
}
