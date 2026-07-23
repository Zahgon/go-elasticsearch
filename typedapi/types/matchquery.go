package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/zerotermsquery"
)

type MatchQuery struct {
	Analyzer *string `json:"analyzer,omitempty"`

	AutoGenerateSynonymsPhraseQuery *bool `json:"auto_generate_synonyms_phrase_query,omitempty"`

	Boost           *float32 `json:"boost,omitempty"`
	CutoffFrequency *Float64 `json:"cutoff_frequency,omitempty"`

	Fuzziness Fuzziness `json:"fuzziness,omitempty"`

	FuzzyRewrite *string `json:"fuzzy_rewrite,omitempty"`

	FuzzyTranspositions *bool `json:"fuzzy_transpositions,omitempty"`

	Lenient *bool `json:"lenient,omitempty"`

	MaxExpansions *int `json:"max_expansions,omitempty"`

	MinimumShouldMatch MinimumShouldMatch `json:"minimum_should_match,omitempty"`

	Operator *operator.Operator `json:"operator,omitempty"`

	PrefixLength *int `json:"prefix_length,omitempty"`

	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`

	ZeroTermsQuery *zerotermsquery.ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

func (s *MatchQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMatchQuery() *MatchQuery { _ = "STUB: not implemented"; return nil }

type MatchQueryVariant interface {
	MatchQueryCaster() *MatchQuery
}

func (s *MatchQuery) MatchQueryCaster() *MatchQuery { _ = "STUB: not implemented"; return nil }
