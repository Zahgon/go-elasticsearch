package types

type AggregationProfileDelegateDebugFilter struct {
	Query                         *string `json:"query,omitempty"`
	ResultsFromMetadata           *int    `json:"results_from_metadata,omitempty"`
	SegmentsCountedInConstantTime *int    `json:"segments_counted_in_constant_time,omitempty"`
	SpecializedFor                *string `json:"specialized_for,omitempty"`
}

func (s *AggregationProfileDelegateDebugFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAggregationProfileDelegateDebugFilter() *AggregationProfileDelegateDebugFilter {
	_ = "STUB: not implemented"
	return nil
}
