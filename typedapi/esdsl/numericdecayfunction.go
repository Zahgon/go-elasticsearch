package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

type _numericDecayFunction struct {
	v *types.NumericDecayFunction
}

func NewNumericDecayFunction() *_numericDecayFunction { _ = "STUB: not implemented"; return nil }

func (s *_numericDecayFunction) DecayFunctionBasedoubledouble(decayfunctionbasedoubledouble map[string]types.DecayPlacementdoubledouble) *_numericDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numericDecayFunction) AddDecayFunctionBasedoubledouble(key string, value types.DecayPlacementdoubledoubleVariant) *_numericDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numericDecayFunction) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *_numericDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numericDecayFunction) FunctionScoreCaster() *types.FunctionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numericDecayFunction) NumericDecayFunctionCaster() *types.NumericDecayFunction {
	_ = "STUB: not implemented"
	return nil
}
