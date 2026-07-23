package types

type ArrayPercentilesItem struct {
	Key           Float64  `json:"key"`
	Value         *Float64 `json:"value,omitempty"`
	ValueAsString *string  `json:"value_as_string,omitempty"`
}

func (s *ArrayPercentilesItem) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewArrayPercentilesItem() *ArrayPercentilesItem { _ = "STUB: not implemented"; return nil }
