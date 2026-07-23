package types

type CartesianBoundsAggregate struct {
	Bounds *TopLeftBottomRightGeoBounds `json:"bounds,omitempty"`
	Meta   Metadata                     `json:"meta,omitempty"`
}

func (s *CartesianBoundsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCartesianBoundsAggregate() *CartesianBoundsAggregate { _ = "STUB: not implemented"; return nil }
