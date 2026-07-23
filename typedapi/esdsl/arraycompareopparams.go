package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/quantifier"
)

type _arrayCompareOpParams struct {
	v *types.ArrayCompareOpParams
}

func NewArrayCompareOpParams(quantifier quantifier.Quantifier) *_arrayCompareOpParams {
	_ = "STUB: not implemented"
	return nil
}

func (s *_arrayCompareOpParams) Quantifier(quantifier quantifier.Quantifier) *_arrayCompareOpParams {
	_ = "STUB: not implemented"
	return nil
}

func (s *_arrayCompareOpParams) Value(fieldvalue types.FieldValueVariant) *_arrayCompareOpParams {
	_ = "STUB: not implemented"
	return nil
}

func (s *_arrayCompareOpParams) ArrayCompareOpParamsCaster() *types.ArrayCompareOpParams {
	_ = "STUB: not implemented"
	return nil
}
