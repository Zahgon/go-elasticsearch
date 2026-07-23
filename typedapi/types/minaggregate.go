package types

type MinAggregate struct {
	Meta Metadata `json:"meta,omitempty"`

	Value         *Float64 `json:"value,omitempty"`
	ValueAsString *string  `json:"value_as_string,omitempty"`
}

func (s *MinAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMinAggregate() *MinAggregate { _ = "STUB: not implemented"; return nil }
