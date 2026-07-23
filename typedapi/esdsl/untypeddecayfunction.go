package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

type _untypedDecayFunction struct {
	v *types.UntypedDecayFunction
}

func NewUntypedDecayFunction() *_untypedDecayFunction { _ = "STUB: not implemented"; return nil }

func (s *_untypedDecayFunction) DecayFunctionBase(decayfunctionbase map[string]types.DecayPlacement) *_untypedDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedDecayFunction) AddDecayFunctionBase(key string, value types.DecayPlacementVariant) *_untypedDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedDecayFunction) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *_untypedDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedDecayFunction) FunctionScoreCaster() *types.FunctionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedDecayFunction) UntypedDecayFunctionCaster() *types.UntypedDecayFunction {
	_ = "STUB: not implemented"
	return nil
}
