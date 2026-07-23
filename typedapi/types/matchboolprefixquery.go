package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
)

type MatchBoolPrefixQuery struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Boost *float32 `json:"boost,omitempty"`

	Fuzziness Fuzziness `json:"fuzziness,omitempty"`

	FuzzyRewrite *string `json:"fuzzy_rewrite,omitempty"`

	FuzzyTranspositions *bool `json:"fuzzy_transpositions,omitempty"`

	MaxExpansions *int `json:"max_expansions,omitempty"`

	MinimumShouldMatch MinimumShouldMatch `json:"minimum_should_match,omitempty"`

	Operator *operator.Operator `json:"operator,omitempty"`

	PrefixLength *int `json:"prefix_length,omitempty"`

	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`
}

func (s *MatchBoolPrefixQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMatchBoolPrefixQuery() *MatchBoolPrefixQuery { _ = "STUB: not implemented"; return nil }

type MatchBoolPrefixQueryVariant interface {
	MatchBoolPrefixQueryCaster() *MatchBoolPrefixQuery
}

func (s *MatchBoolPrefixQuery) MatchBoolPrefixQueryCaster() *MatchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}
