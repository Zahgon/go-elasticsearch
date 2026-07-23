package types

type Intervals struct {
	AllOf *IntervalsAllOf `json:"all_of,omitempty"`

	AnyOf *IntervalsAnyOf `json:"any_of,omitempty"`

	Fuzzy *IntervalsFuzzy `json:"fuzzy,omitempty"`

	Match *IntervalsMatch `json:"match,omitempty"`

	Prefix *IntervalsPrefix `json:"prefix,omitempty"`
	Range  *IntervalsRange  `json:"range,omitempty"`
	Regexp *IntervalsRegexp `json:"regexp,omitempty"`

	Wildcard *IntervalsWildcard `json:"wildcard,omitempty"`
}

func NewIntervals() *Intervals { _ = "STUB: not implemented"; return nil }

type IntervalsVariant interface {
	IntervalsCaster() *Intervals
}

func (s *Intervals) IntervalsCaster() *Intervals { _ = "STUB: not implemented"; return nil }
