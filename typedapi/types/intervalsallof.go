package types

type IntervalsAllOf struct {
	Filter *IntervalsFilter `json:"filter,omitempty"`

	Intervals []Intervals `json:"intervals"`

	MaxGaps *int `json:"max_gaps,omitempty"`

	Ordered *bool `json:"ordered,omitempty"`
}

func (s *IntervalsAllOf) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIntervalsAllOf() *IntervalsAllOf { _ = "STUB: not implemented"; return nil }

type IntervalsAllOfVariant interface {
	IntervalsAllOfCaster() *IntervalsAllOf
}

func (s *IntervalsAllOf) IntervalsAllOfCaster() *IntervalsAllOf {
	_ = "STUB: not implemented"
	return nil
}
