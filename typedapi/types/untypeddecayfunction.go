package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

type UntypedDecayFunction struct {
	DecayFunctionBase map[string]DecayPlacement `json:"-"`

	MultiValueMode *multivaluemode.MultiValueMode `json:"multi_value_mode,omitempty"`
}

func (s *UntypedDecayFunction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s UntypedDecayFunction) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewUntypedDecayFunction() *UntypedDecayFunction { _ = "STUB: not implemented"; return nil }

type UntypedDecayFunctionVariant interface {
	UntypedDecayFunctionCaster() *UntypedDecayFunction
}

func (s *UntypedDecayFunction) UntypedDecayFunctionCaster() *UntypedDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *UntypedDecayFunction) DecayFunctionCaster() *DecayFunction {
	_ = "STUB: not implemented"
	return nil
}
