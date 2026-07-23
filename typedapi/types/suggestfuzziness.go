package types

type SuggestFuzziness struct {
	Fuzziness Fuzziness `json:"fuzziness,omitempty"`

	MinLength *int `json:"min_length,omitempty"`

	PrefixLength *int `json:"prefix_length,omitempty"`

	Transpositions *bool `json:"transpositions,omitempty"`

	UnicodeAware *bool `json:"unicode_aware,omitempty"`
}

func (s *SuggestFuzziness) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSuggestFuzziness() *SuggestFuzziness { _ = "STUB: not implemented"; return nil }

type SuggestFuzzinessVariant interface {
	SuggestFuzzinessCaster() *SuggestFuzziness
}

func (s *SuggestFuzziness) SuggestFuzzinessCaster() *SuggestFuzziness {
	_ = "STUB: not implemented"
	return nil
}
