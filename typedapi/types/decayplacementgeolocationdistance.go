package types

type DecayPlacementGeoLocationDistance struct {
	Decay *Float64 `json:"decay,omitempty"`

	Offset *string `json:"offset,omitempty"`

	Origin GeoLocation `json:"origin,omitempty"`

	Scale *string `json:"scale,omitempty"`
}

func (s *DecayPlacementGeoLocationDistance) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDecayPlacementGeoLocationDistance() *DecayPlacementGeoLocationDistance {
	_ = "STUB: not implemented"
	return nil
}

type DecayPlacementGeoLocationDistanceVariant interface {
	DecayPlacementGeoLocationDistanceCaster() *DecayPlacementGeoLocationDistance
}

func (s *DecayPlacementGeoLocationDistance) DecayPlacementGeoLocationDistanceCaster() *DecayPlacementGeoLocationDistance {
	_ = "STUB: not implemented"
	return nil
}
