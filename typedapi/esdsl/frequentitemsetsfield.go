package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _frequentItemSetsField struct {
	v *types.FrequentItemSetsField
}

func NewFrequentItemSetsField() *_frequentItemSetsField { _ = "STUB: not implemented"; return nil }

func (s *_frequentItemSetsField) Exclude(termsexcludes ...string) *_frequentItemSetsField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_frequentItemSetsField) Field(field string) *_frequentItemSetsField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_frequentItemSetsField) Include(termsinclude types.TermsIncludeVariant) *_frequentItemSetsField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_frequentItemSetsField) FrequentItemSetsFieldCaster() *types.FrequentItemSetsField {
	_ = "STUB: not implemented"
	return nil
}
