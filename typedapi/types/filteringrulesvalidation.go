package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/filteringvalidationstate"
)

type FilteringRulesValidation struct {
	Errors []FilteringValidation                             `json:"errors"`
	State  filteringvalidationstate.FilteringValidationState `json:"state"`
}

func NewFilteringRulesValidation() *FilteringRulesValidation { _ = "STUB: not implemented"; return nil }

type FilteringRulesValidationVariant interface {
	FilteringRulesValidationCaster() *FilteringRulesValidation
}

func (s *FilteringRulesValidation) FilteringRulesValidationCaster() *FilteringRulesValidation {
	_ = "STUB: not implemented"
	return nil
}
