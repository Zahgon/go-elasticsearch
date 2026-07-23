package types

type TopLeftBottomRightGeoBounds struct {
	BottomRight GeoLocation `json:"bottom_right"`
	TopLeft     GeoLocation `json:"top_left"`
}

func (s *TopLeftBottomRightGeoBounds) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTopLeftBottomRightGeoBounds() *TopLeftBottomRightGeoBounds {
	_ = "STUB: not implemented"
	return nil
}

type TopLeftBottomRightGeoBoundsVariant interface {
	TopLeftBottomRightGeoBoundsCaster() *TopLeftBottomRightGeoBounds
}

func (s *TopLeftBottomRightGeoBounds) TopLeftBottomRightGeoBoundsCaster() *TopLeftBottomRightGeoBounds {
	_ = "STUB: not implemented"
	return nil
}

func (s *TopLeftBottomRightGeoBounds) GeoBoundsCaster() *GeoBounds {
	_ = "STUB: not implemented"
	return nil
}
