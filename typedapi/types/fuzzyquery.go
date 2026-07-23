package types

type FuzzyQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Fuzziness Fuzziness `json:"fuzziness,omitempty"`

	MaxExpansions *int `json:"max_expansions,omitempty"`

	PrefixLength *int    `json:"prefix_length,omitempty"`
	QueryName_   *string `json:"_name,omitempty"`

	Rewrite *string `json:"rewrite,omitempty"`

	Transpositions *bool `json:"transpositions,omitempty"`

	Value string `json:"value"`
}

func (s *FuzzyQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFuzzyQuery() *FuzzyQuery { _ = "STUB: not implemented"; return nil }

type FuzzyQueryVariant interface {
	FuzzyQueryCaster() *FuzzyQuery
}

func (s *FuzzyQuery) FuzzyQueryCaster() *FuzzyQuery { _ = "STUB: not implemented"; return nil }
