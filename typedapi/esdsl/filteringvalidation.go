package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _filteringValidation struct {
	v *types.FilteringValidation
}

func NewFilteringValidation() *_filteringValidation { _ = "STUB: not implemented"; return nil }

func (s *_filteringValidation) Ids(ids ...string) *_filteringValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringValidation) Messages(messages ...string) *_filteringValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringValidation) FilteringValidationCaster() *types.FilteringValidation {
	_ = "STUB: not implemented"
	return nil
}
