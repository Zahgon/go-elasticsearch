package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type DerivativeAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
}

func (s *DerivativeAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDerivativeAggregation() *DerivativeAggregation { _ = "STUB: not implemented"; return nil }

type DerivativeAggregationVariant interface {
	DerivativeAggregationCaster() *DerivativeAggregation
}

func (s *DerivativeAggregation) DerivativeAggregationCaster() *DerivativeAggregation {
	_ = "STUB: not implemented"
	return nil
}
