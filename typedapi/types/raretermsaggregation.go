package types

type RareTermsAggregation struct {
	Exclude []string `json:"exclude,omitempty"`

	Field *string `json:"field,omitempty"`

	Include TermsInclude `json:"include,omitempty"`

	MaxDocCount *int64 `json:"max_doc_count,omitempty"`

	Missing Missing `json:"missing,omitempty"`

	Precision *Float64 `json:"precision,omitempty"`
	ValueType *string  `json:"value_type,omitempty"`
}

func (s *RareTermsAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRareTermsAggregation() *RareTermsAggregation { _ = "STUB: not implemented"; return nil }

type RareTermsAggregationVariant interface {
	RareTermsAggregationCaster() *RareTermsAggregation
}

func (s *RareTermsAggregation) RareTermsAggregationCaster() *RareTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}
