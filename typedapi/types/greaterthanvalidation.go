package types

type GreaterThanValidation struct {
	Constraint Float64 `json:"constraint"`
	Type       string  `json:"type,omitempty"`
}

func (s *GreaterThanValidation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s GreaterThanValidation) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGreaterThanValidation() *GreaterThanValidation { _ = "STUB: not implemented"; return nil }

type GreaterThanValidationVariant interface {
	GreaterThanValidationCaster() *GreaterThanValidation
}

func (s *GreaterThanValidation) GreaterThanValidationCaster() *GreaterThanValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *GreaterThanValidation) ValidationCaster() *Validation {
	_ = "STUB: not implemented"
	return nil
}
