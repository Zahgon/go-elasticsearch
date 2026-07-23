package types

type WeightedAvgAggregate struct {
	Meta Metadata `json:"meta,omitempty"`

	Value         *Float64 `json:"value,omitempty"`
	ValueAsString *string  `json:"value_as_string,omitempty"`
}

func (s *WeightedAvgAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewWeightedAvgAggregate() *WeightedAvgAggregate { _ = "STUB: not implemented"; return nil }
