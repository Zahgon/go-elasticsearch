package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _greaterThanValidation struct {
	v *types.GreaterThanValidation
}

func NewGreaterThanValidation(constraint types.Float64) *_greaterThanValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_greaterThanValidation) Constraint(constraint types.Float64) *_greaterThanValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_greaterThanValidation) GreaterThanValidationCaster() *types.GreaterThanValidation {
	_ = "STUB: not implemented"
	return nil
}
