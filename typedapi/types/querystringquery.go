package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/textquerytype"
)

type QueryStringQuery struct {
	AllowLeadingWildcard *bool `json:"allow_leading_wildcard,omitempty"`

	AnalyzeWildcard *bool `json:"analyze_wildcard,omitempty"`

	Analyzer *string `json:"analyzer,omitempty"`

	AutoGenerateSynonymsPhraseQuery *bool `json:"auto_generate_synonyms_phrase_query,omitempty"`

	Boost *float32 `json:"boost,omitempty"`

	DefaultField *string `json:"default_field,omitempty"`

	DefaultOperator *operator.Operator `json:"default_operator,omitempty"`

	EnablePositionIncrements *bool `json:"enable_position_increments,omitempty"`
	Escape                   *bool `json:"escape,omitempty"`

	Fields []string `json:"fields,omitempty"`

	Fuzziness Fuzziness `json:"fuzziness,omitempty"`

	FuzzyMaxExpansions *int `json:"fuzzy_max_expansions,omitempty"`

	FuzzyPrefixLength *int `json:"fuzzy_prefix_length,omitempty"`

	FuzzyRewrite *string `json:"fuzzy_rewrite,omitempty"`

	FuzzyTranspositions *bool `json:"fuzzy_transpositions,omitempty"`

	Lenient *bool `json:"lenient,omitempty"`

	MaxDeterminizedStates *int `json:"max_determinized_states,omitempty"`

	MinimumShouldMatch MinimumShouldMatch `json:"minimum_should_match,omitempty"`

	PhraseSlop *Float64 `json:"phrase_slop,omitempty"`

	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`

	QuoteAnalyzer *string `json:"quote_analyzer,omitempty"`

	QuoteFieldSuffix *string `json:"quote_field_suffix,omitempty"`

	Rewrite *string `json:"rewrite,omitempty"`

	TieBreaker *Float64 `json:"tie_breaker,omitempty"`

	TimeZone *string `json:"time_zone,omitempty"`

	Type *textquerytype.TextQueryType `json:"type,omitempty"`
}

func (s *QueryStringQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewQueryStringQuery() *QueryStringQuery { _ = "STUB: not implemented"; return nil }

type QueryStringQueryVariant interface {
	QueryStringQueryCaster() *QueryStringQuery
}

func (s *QueryStringQuery) QueryStringQueryCaster() *QueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}
