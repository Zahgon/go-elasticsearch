package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _lessThanValidation struct {
	v *types.LessThanValidation
}

func NewLessThanValidation(constraint types.Float64) *_lessThanValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lessThanValidation) Constraint(constraint types.Float64) *_lessThanValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lessThanValidation) LessThanValidationCaster() *types.LessThanValidation {
	_ = "STUB: not implemented"
	return nil
}
