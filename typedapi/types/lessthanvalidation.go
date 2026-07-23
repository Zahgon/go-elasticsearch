package types

type LessThanValidation struct {
	Constraint Float64 `json:"constraint"`
	Type       string  `json:"type,omitempty"`
}

func (s *LessThanValidation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s LessThanValidation) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLessThanValidation() *LessThanValidation { _ = "STUB: not implemented"; return nil }

type LessThanValidationVariant interface {
	LessThanValidationCaster() *LessThanValidation
}

func (s *LessThanValidation) LessThanValidationCaster() *LessThanValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *LessThanValidation) ValidationCaster() *Validation { _ = "STUB: not implemented"; return nil }
