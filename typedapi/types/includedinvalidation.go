package types

type IncludedInValidation struct {
	Constraint []ScalarValue `json:"constraint"`
	Type       string        `json:"type,omitempty"`
}

func (s IncludedInValidation) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewIncludedInValidation() *IncludedInValidation { _ = "STUB: not implemented"; return nil }

type IncludedInValidationVariant interface {
	IncludedInValidationCaster() *IncludedInValidation
}

func (s *IncludedInValidation) IncludedInValidationCaster() *IncludedInValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *IncludedInValidation) ValidationCaster() *Validation {
	_ = "STUB: not implemented"
	return nil
}
