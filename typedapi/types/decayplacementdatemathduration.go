package types

type DecayPlacementDateMathDuration struct {
	Decay *Float64 `json:"decay,omitempty"`

	Offset Duration `json:"offset,omitempty"`

	Origin *string `json:"origin,omitempty"`

	Scale Duration `json:"scale,omitempty"`
}

func (s *DecayPlacementDateMathDuration) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDecayPlacementDateMathDuration() *DecayPlacementDateMathDuration {
	_ = "STUB: not implemented"
	return nil
}

type DecayPlacementDateMathDurationVariant interface {
	DecayPlacementDateMathDurationCaster() *DecayPlacementDateMathDuration
}

func (s *DecayPlacementDateMathDuration) DecayPlacementDateMathDurationCaster() *DecayPlacementDateMathDuration {
	_ = "STUB: not implemented"
	return nil
}
