package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _includedInValidation struct {
	v *types.IncludedInValidation
}

func NewIncludedInValidation() *_includedInValidation { _ = "STUB: not implemented"; return nil }

func (s *_includedInValidation) Constraint(constraints ...types.ScalarValueVariant) *_includedInValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_includedInValidation) ConstraintValues(constraintvalues []types.ScalarValue) *_includedInValidation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_includedInValidation) IncludedInValidationCaster() *types.IncludedInValidation {
	_ = "STUB: not implemented"
	return nil
}
