package types

type MaxAggregate struct {
	Meta Metadata `json:"meta,omitempty"`

	Value         *Float64 `json:"value,omitempty"`
	ValueAsString *string  `json:"value_as_string,omitempty"`
}

func (s *MaxAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMaxAggregate() *MaxAggregate { _ = "STUB: not implemented"; return nil }
