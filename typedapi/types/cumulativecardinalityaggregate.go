package types

type CumulativeCardinalityAggregate struct {
	Meta          Metadata `json:"meta,omitempty"`
	Value         int64    `json:"value"`
	ValueAsString *string  `json:"value_as_string,omitempty"`
}

func (s *CumulativeCardinalityAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCumulativeCardinalityAggregate() *CumulativeCardinalityAggregate {
	_ = "STUB: not implemented"
	return nil
}
