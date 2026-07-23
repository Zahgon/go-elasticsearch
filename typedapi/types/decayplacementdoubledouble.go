package types

type DecayPlacementdoubledouble struct {
	Decay *Float64 `json:"decay,omitempty"`

	Offset *Float64 `json:"offset,omitempty"`

	Origin *Float64 `json:"origin,omitempty"`

	Scale *Float64 `json:"scale,omitempty"`
}

func (s *DecayPlacementdoubledouble) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDecayPlacementdoubledouble() *DecayPlacementdoubledouble {
	_ = "STUB: not implemented"
	return nil
}

type DecayPlacementdoubledoubleVariant interface {
	DecayPlacementdoubledoubleCaster() *DecayPlacementdoubledouble
}

func (s *DecayPlacementdoubledouble) DecayPlacementdoubledoubleCaster() *DecayPlacementdoubledouble {
	_ = "STUB: not implemented"
	return nil
}
