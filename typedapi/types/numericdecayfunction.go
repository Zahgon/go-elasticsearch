package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

type NumericDecayFunction struct {
	DecayFunctionBasedoubledouble map[string]DecayPlacementdoubledouble `json:"-"`

	MultiValueMode *multivaluemode.MultiValueMode `json:"multi_value_mode,omitempty"`
}

func (s *NumericDecayFunction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s NumericDecayFunction) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewNumericDecayFunction() *NumericDecayFunction { _ = "STUB: not implemented"; return nil }

type NumericDecayFunctionVariant interface {
	NumericDecayFunctionCaster() *NumericDecayFunction
}

func (s *NumericDecayFunction) NumericDecayFunctionCaster() *NumericDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *NumericDecayFunction) DecayFunctionCaster() *DecayFunction {
	_ = "STUB: not implemented"
	return nil
}
