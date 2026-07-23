package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type CumulativeSumAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
}

func (s *CumulativeSumAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCumulativeSumAggregation() *CumulativeSumAggregation { _ = "STUB: not implemented"; return nil }

type CumulativeSumAggregationVariant interface {
	CumulativeSumAggregationCaster() *CumulativeSumAggregation
}

func (s *CumulativeSumAggregation) CumulativeSumAggregationCaster() *CumulativeSumAggregation {
	_ = "STUB: not implemented"
	return nil
}
