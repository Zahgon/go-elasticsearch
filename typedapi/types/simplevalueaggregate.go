package types

type SimpleValueAggregate struct {
	Meta Metadata `json:"meta,omitempty"`

	Value         *Float64 `json:"value,omitempty"`
	ValueAsString *string  `json:"value_as_string,omitempty"`
}

func (s *SimpleValueAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSimpleValueAggregate() *SimpleValueAggregate { _ = "STUB: not implemented"; return nil }
