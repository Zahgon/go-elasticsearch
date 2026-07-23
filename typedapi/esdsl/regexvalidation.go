package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _regexValidation struct {
	v *types.RegexValidation
}

func NewRegexValidation(constraint string) *_regexValidation { _ = "STUB: not implemented"; return nil }

func (s *_regexValidation) Constraint(constraint string) *_regexValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_regexValidation) RegexValidationCaster() *types.RegexValidation {
	_ = "STUB: not implemented"
	return nil
}
