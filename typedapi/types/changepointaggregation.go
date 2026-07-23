package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type ChangePointAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
}

func (s *ChangePointAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewChangePointAggregation() *ChangePointAggregation { _ = "STUB: not implemented"; return nil }

type ChangePointAggregationVariant interface {
	ChangePointAggregationCaster() *ChangePointAggregation
}

func (s *ChangePointAggregation) ChangePointAggregationCaster() *ChangePointAggregation {
	_ = "STUB: not implemented"
	return nil
}
