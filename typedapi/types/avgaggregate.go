package types

type AvgAggregate struct {
	Meta Metadata `json:"meta,omitempty"`

	Value         *Float64 `json:"value,omitempty"`
	ValueAsString *string  `json:"value_as_string,omitempty"`
}

func (s *AvgAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAvgAggregate() *AvgAggregate { _ = "STUB: not implemented"; return nil }
