package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _decayFunction struct {
	v types.DecayFunction
}

func NewDecayFunction() *_decayFunction { _ = "STUB: not implemented"; return nil }

func (u *_decayFunction) UntypedDecayFunction(untypeddecayfunction types.UntypedDecayFunctionVariant) *_decayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (u *_untypedDecayFunction) DecayFunctionCaster() *types.DecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (u *_decayFunction) DateDecayFunction(datedecayfunction types.DateDecayFunctionVariant) *_decayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (u *_dateDecayFunction) DecayFunctionCaster() *types.DecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (u *_decayFunction) NumericDecayFunction(numericdecayfunction types.NumericDecayFunctionVariant) *_decayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (u *_numericDecayFunction) DecayFunctionCaster() *types.DecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (u *_decayFunction) GeoDecayFunction(geodecayfunction types.GeoDecayFunctionVariant) *_decayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (u *_geoDecayFunction) DecayFunctionCaster() *types.DecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (u *_decayFunction) DecayFunctionCaster() *types.DecayFunction {
	_ = "STUB: not implemented"
	return nil
}
