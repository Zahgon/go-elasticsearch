package types

type DerivativeAggregate struct {
	Meta                    Metadata `json:"meta,omitempty"`
	NormalizedValue         *Float64 `json:"normalized_value,omitempty"`
	NormalizedValueAsString *string  `json:"normalized_value_as_string,omitempty"`

	Value         *Float64 `json:"value,omitempty"`
	ValueAsString *string  `json:"value_as_string,omitempty"`
}

func (s *DerivativeAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDerivativeAggregate() *DerivativeAggregate { _ = "STUB: not implemented"; return nil }
