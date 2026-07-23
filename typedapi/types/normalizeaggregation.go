package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/normalizemethod"
)

type NormalizeAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`

	Method *normalizemethod.NormalizeMethod `json:"method,omitempty"`
}

func (s *NormalizeAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNormalizeAggregation() *NormalizeAggregation { _ = "STUB: not implemented"; return nil }

type NormalizeAggregationVariant interface {
	NormalizeAggregationCaster() *NormalizeAggregation
}

func (s *NormalizeAggregation) NormalizeAggregationCaster() *NormalizeAggregation {
	_ = "STUB: not implemented"
	return nil
}
