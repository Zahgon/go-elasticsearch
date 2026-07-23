package types

type IntervalsAnyOf struct {
	Filter *IntervalsFilter `json:"filter,omitempty"`

	Intervals []Intervals `json:"intervals"`
}

func NewIntervalsAnyOf() *IntervalsAnyOf { _ = "STUB: not implemented"; return nil }

type IntervalsAnyOfVariant interface {
	IntervalsAnyOfCaster() *IntervalsAnyOf
}

func (s *IntervalsAnyOf) IntervalsAnyOfCaster() *IntervalsAnyOf {
	_ = "STUB: not implemented"
	return nil
}
