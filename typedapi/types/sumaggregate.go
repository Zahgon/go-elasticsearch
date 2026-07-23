package types

type SumAggregate struct {
	Meta Metadata `json:"meta,omitempty"`

	Value         *Float64 `json:"value,omitempty"`
	ValueAsString *string  `json:"value_as_string,omitempty"`
}

func (s *SumAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSumAggregate() *SumAggregate { _ = "STUB: not implemented"; return nil }
