package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type BucketScriptAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`

	Script *Script `json:"script,omitempty"`
}

func (s *BucketScriptAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewBucketScriptAggregation() *BucketScriptAggregation { _ = "STUB: not implemented"; return nil }

type BucketScriptAggregationVariant interface {
	BucketScriptAggregationCaster() *BucketScriptAggregation
}

func (s *BucketScriptAggregation) BucketScriptAggregationCaster() *BucketScriptAggregation {
	_ = "STUB: not implemented"
	return nil
}
