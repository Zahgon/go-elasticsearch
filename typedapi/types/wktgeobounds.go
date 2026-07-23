package types

type WktGeoBounds struct {
	Wkt string `json:"wkt"`
}

func (s *WktGeoBounds) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWktGeoBounds() *WktGeoBounds { _ = "STUB: not implemented"; return nil }

type WktGeoBoundsVariant interface {
	WktGeoBoundsCaster() *WktGeoBounds
}

func (s *WktGeoBounds) WktGeoBoundsCaster() *WktGeoBounds { _ = "STUB: not implemented"; return nil }

func (s *WktGeoBounds) GeoBoundsCaster() *GeoBounds { _ = "STUB: not implemented"; return nil }
