package types

type IntervalsFuzzy struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Fuzziness Fuzziness `json:"fuzziness,omitempty"`

	PrefixLength *int `json:"prefix_length,omitempty"`

	Term string `json:"term"`

	Transpositions *bool `json:"transpositions,omitempty"`

	UseField *string `json:"use_field,omitempty"`
}

func (s *IntervalsFuzzy) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIntervalsFuzzy() *IntervalsFuzzy { _ = "STUB: not implemented"; return nil }

type IntervalsFuzzyVariant interface {
	IntervalsFuzzyCaster() *IntervalsFuzzy
}

func (s *IntervalsFuzzy) IntervalsFuzzyCaster() *IntervalsFuzzy {
	_ = "STUB: not implemented"
	return nil
}
