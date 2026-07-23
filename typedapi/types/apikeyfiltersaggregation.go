package types

type ApiKeyFiltersAggregation struct {
	Filters BucketsApiKeyQueryContainer `json:"filters,omitempty"`

	Keyed *bool `json:"keyed,omitempty"`

	OtherBucket *bool `json:"other_bucket,omitempty"`

	OtherBucketKey *string `json:"other_bucket_key,omitempty"`
}

func (s *ApiKeyFiltersAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewApiKeyFiltersAggregation() *ApiKeyFiltersAggregation { _ = "STUB: not implemented"; return nil }

type ApiKeyFiltersAggregationVariant interface {
	ApiKeyFiltersAggregationCaster() *ApiKeyFiltersAggregation
}

func (s *ApiKeyFiltersAggregation) ApiKeyFiltersAggregationCaster() *ApiKeyFiltersAggregation {
	_ = "STUB: not implemented"
	return nil
}
