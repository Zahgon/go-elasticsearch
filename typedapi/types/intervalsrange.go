package types

type IntervalsRange struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Gt *string `json:"gt,omitempty"`

	Gte *string `json:"gte,omitempty"`

	Lt *string `json:"lt,omitempty"`

	Lte *string `json:"lte,omitempty"`

	UseField *string `json:"use_field,omitempty"`
}

func (s *IntervalsRange) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIntervalsRange() *IntervalsRange { _ = "STUB: not implemented"; return nil }

type IntervalsRangeVariant interface {
	IntervalsRangeCaster() *IntervalsRange
}

func (s *IntervalsRange) IntervalsRangeCaster() *IntervalsRange {
	_ = "STUB: not implemented"
	return nil
}
