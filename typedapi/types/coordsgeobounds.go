package types

type CoordsGeoBounds struct {
	Bottom Float64 `json:"bottom"`
	Left   Float64 `json:"left"`
	Right  Float64 `json:"right"`
	Top    Float64 `json:"top"`
}

func (s *CoordsGeoBounds) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCoordsGeoBounds() *CoordsGeoBounds { _ = "STUB: not implemented"; return nil }

type CoordsGeoBoundsVariant interface {
	CoordsGeoBoundsCaster() *CoordsGeoBounds
}

func (s *CoordsGeoBounds) CoordsGeoBoundsCaster() *CoordsGeoBounds {
	_ = "STUB: not implemented"
	return nil
}

func (s *CoordsGeoBounds) GeoBoundsCaster() *GeoBounds { _ = "STUB: not implemented"; return nil }
