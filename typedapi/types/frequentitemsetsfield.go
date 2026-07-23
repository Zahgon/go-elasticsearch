package types

type FrequentItemSetsField struct {
	Exclude []string `json:"exclude,omitempty"`
	Field   string   `json:"field"`

	Include TermsInclude `json:"include,omitempty"`
}

func (s *FrequentItemSetsField) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFrequentItemSetsField() *FrequentItemSetsField { _ = "STUB: not implemented"; return nil }

type FrequentItemSetsFieldVariant interface {
	FrequentItemSetsFieldCaster() *FrequentItemSetsField
}

func (s *FrequentItemSetsField) FrequentItemSetsFieldCaster() *FrequentItemSetsField {
	_ = "STUB: not implemented"
	return nil
}
