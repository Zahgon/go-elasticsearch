package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
)

type SimpleQueryStringQuery struct {
	AnalyzeWildcard *bool `json:"analyze_wildcard,omitempty"`

	Analyzer *string `json:"analyzer,omitempty"`

	AutoGenerateSynonymsPhraseQuery *bool `json:"auto_generate_synonyms_phrase_query,omitempty"`

	Boost *float32 `json:"boost,omitempty"`

	DefaultOperator *operator.Operator `json:"default_operator,omitempty"`

	Fields []string `json:"fields,omitempty"`

	Flags PipeSeparatedFlagsSimpleQueryStringFlag `json:"flags,omitempty"`

	FuzzyMaxExpansions *int `json:"fuzzy_max_expansions,omitempty"`

	FuzzyPrefixLength *int `json:"fuzzy_prefix_length,omitempty"`

	FuzzyTranspositions *bool `json:"fuzzy_transpositions,omitempty"`

	Lenient *bool `json:"lenient,omitempty"`

	MinimumShouldMatch MinimumShouldMatch `json:"minimum_should_match,omitempty"`

	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`

	QuoteFieldSuffix *string `json:"quote_field_suffix,omitempty"`
}

func (s *SimpleQueryStringQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSimpleQueryStringQuery() *SimpleQueryStringQuery { _ = "STUB: not implemented"; return nil }

type SimpleQueryStringQueryVariant interface {
	SimpleQueryStringQueryCaster() *SimpleQueryStringQuery
}

func (s *SimpleQueryStringQuery) SimpleQueryStringQueryCaster() *SimpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}
