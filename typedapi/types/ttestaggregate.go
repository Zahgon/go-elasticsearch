package types

type TTestAggregate struct {
	Meta          Metadata `json:"meta,omitempty"`
	Value         *Float64 `json:"value,omitempty"`
	ValueAsString *string  `json:"value_as_string,omitempty"`
}

func (s *TTestAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTTestAggregate() *TTestAggregate { _ = "STUB: not implemented"; return nil }
