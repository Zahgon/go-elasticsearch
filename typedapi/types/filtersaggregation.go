package types

type FiltersAggregation struct {
	Filters BucketsQuery `json:"filters,omitempty"`

	Keyed *bool `json:"keyed,omitempty"`

	OtherBucket *bool `json:"other_bucket,omitempty"`

	OtherBucketKey *string `json:"other_bucket_key,omitempty"`
}

func (s *FiltersAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFiltersAggregation() *FiltersAggregation { _ = "STUB: not implemented"; return nil }

type FiltersAggregationVariant interface {
	FiltersAggregationCaster() *FiltersAggregation
}

func (s *FiltersAggregation) FiltersAggregationCaster() *FiltersAggregation {
	_ = "STUB: not implemented"
	return nil
}
