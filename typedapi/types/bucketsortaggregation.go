package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type BucketSortAggregation struct {
	From *int `json:"from,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`

	Size *int `json:"size,omitempty"`

	Sort []SortCombinations `json:"sort,omitempty"`
}

func (s *BucketSortAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewBucketSortAggregation() *BucketSortAggregation { _ = "STUB: not implemented"; return nil }

type BucketSortAggregationVariant interface {
	BucketSortAggregationCaster() *BucketSortAggregation
}

func (s *BucketSortAggregation) BucketSortAggregationCaster() *BucketSortAggregation {
	_ = "STUB: not implemented"
	return nil
}
