package types

type IntervalsMatch struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Filter *IntervalsFilter `json:"filter,omitempty"`

	MaxGaps *int `json:"max_gaps,omitempty"`

	Ordered *bool `json:"ordered,omitempty"`

	Query string `json:"query"`

	UseField *string `json:"use_field,omitempty"`
}

func (s *IntervalsMatch) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIntervalsMatch() *IntervalsMatch { _ = "STUB: not implemented"; return nil }

type IntervalsMatchVariant interface {
	IntervalsMatchCaster() *IntervalsMatch
}

func (s *IntervalsMatch) IntervalsMatchCaster() *IntervalsMatch {
	_ = "STUB: not implemented"
	return nil
}
