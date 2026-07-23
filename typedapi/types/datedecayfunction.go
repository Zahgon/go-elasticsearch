package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

type DateDecayFunction struct {
	DecayFunctionBaseDateMathDuration map[string]DecayPlacementDateMathDuration `json:"-"`

	MultiValueMode *multivaluemode.MultiValueMode `json:"multi_value_mode,omitempty"`
}

func (s *DateDecayFunction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s DateDecayFunction) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDateDecayFunction() *DateDecayFunction { _ = "STUB: not implemented"; return nil }

type DateDecayFunctionVariant interface {
	DateDecayFunctionCaster() *DateDecayFunction
}

func (s *DateDecayFunction) DateDecayFunctionCaster() *DateDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *DateDecayFunction) DecayFunctionCaster() *DecayFunction {
	_ = "STUB: not implemented"
	return nil
}
