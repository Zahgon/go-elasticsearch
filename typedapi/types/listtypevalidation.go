package types

type ListTypeValidation struct {
	Constraint string `json:"constraint"`
	Type       string `json:"type,omitempty"`
}

func (s *ListTypeValidation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ListTypeValidation) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewListTypeValidation() *ListTypeValidation { _ = "STUB: not implemented"; return nil }

type ListTypeValidationVariant interface {
	ListTypeValidationCaster() *ListTypeValidation
}

func (s *ListTypeValidation) ListTypeValidationCaster() *ListTypeValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *ListTypeValidation) ValidationCaster() *Validation { _ = "STUB: not implemented"; return nil }
