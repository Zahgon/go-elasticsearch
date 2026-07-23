package types

type IntervalsWildcard struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Pattern string `json:"pattern"`

	UseField *string `json:"use_field,omitempty"`
}

func (s *IntervalsWildcard) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIntervalsWildcard() *IntervalsWildcard { _ = "STUB: not implemented"; return nil }

type IntervalsWildcardVariant interface {
	IntervalsWildcardCaster() *IntervalsWildcard
}

func (s *IntervalsWildcard) IntervalsWildcardCaster() *IntervalsWildcard {
	_ = "STUB: not implemented"
	return nil
}
