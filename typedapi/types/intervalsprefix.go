package types

type IntervalsPrefix struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Prefix string `json:"prefix"`

	UseField *string `json:"use_field,omitempty"`
}

func (s *IntervalsPrefix) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIntervalsPrefix() *IntervalsPrefix { _ = "STUB: not implemented"; return nil }

type IntervalsPrefixVariant interface {
	IntervalsPrefixCaster() *IntervalsPrefix
}

func (s *IntervalsPrefix) IntervalsPrefixCaster() *IntervalsPrefix {
	_ = "STUB: not implemented"
	return nil
}
