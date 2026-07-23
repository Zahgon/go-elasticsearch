package types

type IntervalsQuery struct {
	AllOf *IntervalsAllOf `json:"all_of,omitempty"`

	AnyOf *IntervalsAnyOf `json:"any_of,omitempty"`

	Boost *float32 `json:"boost,omitempty"`

	Fuzzy *IntervalsFuzzy `json:"fuzzy,omitempty"`

	Match *IntervalsMatch `json:"match,omitempty"`

	Prefix     *IntervalsPrefix `json:"prefix,omitempty"`
	QueryName_ *string          `json:"_name,omitempty"`
	Range      *IntervalsRange  `json:"range,omitempty"`
	Regexp     *IntervalsRegexp `json:"regexp,omitempty"`

	Wildcard *IntervalsWildcard `json:"wildcard,omitempty"`
}

func (s *IntervalsQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIntervalsQuery() *IntervalsQuery { _ = "STUB: not implemented"; return nil }

type IntervalsQueryVariant interface {
	IntervalsQueryCaster() *IntervalsQuery
}

func (s *IntervalsQuery) IntervalsQueryCaster() *IntervalsQuery {
	_ = "STUB: not implemented"
	return nil
}
