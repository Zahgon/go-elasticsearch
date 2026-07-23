package types

type RegexValidation struct {
	Constraint string `json:"constraint"`
	Type       string `json:"type,omitempty"`
}

func (s *RegexValidation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s RegexValidation) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewRegexValidation() *RegexValidation { _ = "STUB: not implemented"; return nil }

type RegexValidationVariant interface {
	RegexValidationCaster() *RegexValidation
}

func (s *RegexValidation) RegexValidationCaster() *RegexValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *RegexValidation) ValidationCaster() *Validation { _ = "STUB: not implemented"; return nil }
