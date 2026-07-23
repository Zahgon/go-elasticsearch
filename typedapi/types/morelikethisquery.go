package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type MoreLikeThisQuery struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Boost *float32 `json:"boost,omitempty"`

	BoostTerms *Float64 `json:"boost_terms,omitempty"`

	FailOnUnsupportedField *bool `json:"fail_on_unsupported_field,omitempty"`

	Fields []string `json:"fields,omitempty"`

	Include *bool `json:"include,omitempty"`

	Like []Like `json:"like"`

	MaxDocFreq *int `json:"max_doc_freq,omitempty"`

	MaxQueryTerms *int `json:"max_query_terms,omitempty"`

	MaxWordLength *int `json:"max_word_length,omitempty"`

	MinDocFreq *int `json:"min_doc_freq,omitempty"`

	MinTermFreq *int `json:"min_term_freq,omitempty"`

	MinWordLength *int `json:"min_word_length,omitempty"`

	MinimumShouldMatch MinimumShouldMatch `json:"minimum_should_match,omitempty"`
	QueryName_         *string            `json:"_name,omitempty"`
	Routing            *string            `json:"routing,omitempty"`

	StopWords StopWords `json:"stop_words,omitempty"`

	Unlike      []Like                   `json:"unlike,omitempty"`
	Version     *int64                   `json:"version,omitempty"`
	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *MoreLikeThisQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMoreLikeThisQuery() *MoreLikeThisQuery { _ = "STUB: not implemented"; return nil }

type MoreLikeThisQueryVariant interface {
	MoreLikeThisQueryCaster() *MoreLikeThisQuery
}

func (s *MoreLikeThisQuery) MoreLikeThisQueryCaster() *MoreLikeThisQuery {
	_ = "STUB: not implemented"
	return nil
}
