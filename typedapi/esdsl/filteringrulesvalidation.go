package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/filteringvalidationstate"
)

type _filteringRulesValidation struct {
	v *types.FilteringRulesValidation
}

func NewFilteringRulesValidation(state filteringvalidationstate.FilteringValidationState) *_filteringRulesValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRulesValidation) Errors(errors ...types.FilteringValidationVariant) *_filteringRulesValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRulesValidation) ErrorsValues(errorsvalues []types.FilteringValidation) *_filteringRulesValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRulesValidation) State(state filteringvalidationstate.FilteringValidationState) *_filteringRulesValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRulesValidation) FilteringRulesValidationCaster() *types.FilteringRulesValidation {
	_ = "STUB: not implemented"
	return nil
}
