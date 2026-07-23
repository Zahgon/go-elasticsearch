package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

type _dateDecayFunction struct {
	v *types.DateDecayFunction
}

func NewDateDecayFunction() *_dateDecayFunction { _ = "STUB: not implemented"; return nil }

func (s *_dateDecayFunction) DecayFunctionBaseDateMathDuration(decayfunctionbasedatemathduration map[string]types.DecayPlacementDateMathDuration) *_dateDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateDecayFunction) AddDecayFunctionBaseDateMathDuration(key string, value types.DecayPlacementDateMathDurationVariant) *_dateDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateDecayFunction) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *_dateDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateDecayFunction) FunctionScoreCaster() *types.FunctionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateDecayFunction) DateDecayFunctionCaster() *types.DateDecayFunction {
	_ = "STUB: not implemented"
	return nil
}
