package types

type GeoBoundsAggregate struct {
	Bounds GeoBounds `json:"bounds,omitempty"`
	Meta   Metadata  `json:"meta,omitempty"`
}

func (s *GeoBoundsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoBoundsAggregate() *GeoBoundsAggregate { _ = "STUB: not implemented"; return nil }
