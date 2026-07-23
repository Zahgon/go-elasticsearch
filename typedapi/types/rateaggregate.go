package types

type RateAggregate struct {
	Meta          Metadata `json:"meta,omitempty"`
	Value         Float64  `json:"value"`
	ValueAsString *string  `json:"value_as_string,omitempty"`
}

func (s *RateAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRateAggregate() *RateAggregate { _ = "STUB: not implemented"; return nil }
