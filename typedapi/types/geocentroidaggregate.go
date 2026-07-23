package types

type GeoCentroidAggregate struct {
	Count    int64       `json:"count"`
	Location GeoLocation `json:"location,omitempty"`
	Meta     Metadata    `json:"meta,omitempty"`
}

func (s *GeoCentroidAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoCentroidAggregate() *GeoCentroidAggregate { _ = "STUB: not implemented"; return nil }
