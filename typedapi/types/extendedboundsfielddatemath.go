package types

type ExtendedBoundsFieldDateMath struct {
	Max FieldDateMath `json:"max,omitempty"`

	Min FieldDateMath `json:"min,omitempty"`
}

func (s *ExtendedBoundsFieldDateMath) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewExtendedBoundsFieldDateMath() *ExtendedBoundsFieldDateMath {
	_ = "STUB: not implemented"
	return nil
}

type ExtendedBoundsFieldDateMathVariant interface {
	ExtendedBoundsFieldDateMathCaster() *ExtendedBoundsFieldDateMath
}

func (s *ExtendedBoundsFieldDateMath) ExtendedBoundsFieldDateMathCaster() *ExtendedBoundsFieldDateMath {
	_ = "STUB: not implemented"
	return nil
}
