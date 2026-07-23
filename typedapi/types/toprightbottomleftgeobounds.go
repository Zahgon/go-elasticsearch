package types

type TopRightBottomLeftGeoBounds struct {
	BottomLeft GeoLocation `json:"bottom_left"`
	TopRight   GeoLocation `json:"top_right"`
}

func (s *TopRightBottomLeftGeoBounds) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTopRightBottomLeftGeoBounds() *TopRightBottomLeftGeoBounds {
	_ = "STUB: not implemented"
	return nil
}

type TopRightBottomLeftGeoBoundsVariant interface {
	TopRightBottomLeftGeoBoundsCaster() *TopRightBottomLeftGeoBounds
}

func (s *TopRightBottomLeftGeoBounds) TopRightBottomLeftGeoBoundsCaster() *TopRightBottomLeftGeoBounds {
	_ = "STUB: not implemented"
	return nil
}

func (s *TopRightBottomLeftGeoBounds) GeoBoundsCaster() *GeoBounds {
	_ = "STUB: not implemented"
	return nil
}
