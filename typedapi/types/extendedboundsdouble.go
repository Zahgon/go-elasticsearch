package types

type ExtendedBoundsdouble struct {
	Max *Float64 `json:"max,omitempty"`

	Min *Float64 `json:"min,omitempty"`
}

func (s *ExtendedBoundsdouble) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewExtendedBoundsdouble() *ExtendedBoundsdouble { _ = "STUB: not implemented"; return nil }

type ExtendedBoundsdoubleVariant interface {
	ExtendedBoundsdoubleCaster() *ExtendedBoundsdouble
}

func (s *ExtendedBoundsdouble) ExtendedBoundsdoubleCaster() *ExtendedBoundsdouble {
	_ = "STUB: not implemented"
	return nil
}
