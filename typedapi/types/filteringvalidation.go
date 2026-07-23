package types

type FilteringValidation struct {
	Ids      []string `json:"ids"`
	Messages []string `json:"messages"`
}

func NewFilteringValidation() *FilteringValidation { _ = "STUB: not implemented"; return nil }

type FilteringValidationVariant interface {
	FilteringValidationCaster() *FilteringValidation
}

func (s *FilteringValidation) FilteringValidationCaster() *FilteringValidation {
	_ = "STUB: not implemented"
	return nil
}
