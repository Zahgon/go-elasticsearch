package types

type CartesianCentroidAggregate struct {
	Count    int64           `json:"count"`
	Location *CartesianPoint `json:"location,omitempty"`
	Meta     Metadata        `json:"meta,omitempty"`
}

func (s *CartesianCentroidAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCartesianCentroidAggregate() *CartesianCentroidAggregate {
	_ = "STUB: not implemented"
	return nil
}
