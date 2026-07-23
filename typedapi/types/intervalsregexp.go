package types

type IntervalsRegexp struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Pattern string `json:"pattern"`

	UseField *string `json:"use_field,omitempty"`
}

func (s *IntervalsRegexp) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIntervalsRegexp() *IntervalsRegexp { _ = "STUB: not implemented"; return nil }

type IntervalsRegexpVariant interface {
	IntervalsRegexpCaster() *IntervalsRegexp
}

func (s *IntervalsRegexp) IntervalsRegexpCaster() *IntervalsRegexp {
	_ = "STUB: not implemented"
	return nil
}
