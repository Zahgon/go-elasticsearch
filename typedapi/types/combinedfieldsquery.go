package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/combinedfieldsoperator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/combinedfieldszeroterms"
)

type CombinedFieldsQuery struct {
	AutoGenerateSynonymsPhraseQuery *bool `json:"auto_generate_synonyms_phrase_query,omitempty"`

	Boost *float32 `json:"boost,omitempty"`

	Fields []string `json:"fields"`

	MinimumShouldMatch MinimumShouldMatch `json:"minimum_should_match,omitempty"`

	Operator *combinedfieldsoperator.CombinedFieldsOperator `json:"operator,omitempty"`

	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`

	ZeroTermsQuery *combinedfieldszeroterms.CombinedFieldsZeroTerms `json:"zero_terms_query,omitempty"`
}

func (s *CombinedFieldsQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCombinedFieldsQuery() *CombinedFieldsQuery { _ = "STUB: not implemented"; return nil }

type CombinedFieldsQueryVariant interface {
	CombinedFieldsQueryCaster() *CombinedFieldsQuery
}

func (s *CombinedFieldsQuery) CombinedFieldsQueryCaster() *CombinedFieldsQuery {
	_ = "STUB: not implemented"
	return nil
}
