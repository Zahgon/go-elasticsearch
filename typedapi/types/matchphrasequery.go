package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/zerotermsquery"
)

type MatchPhraseQuery struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Boost *float32 `json:"boost,omitempty"`

	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`

	Slop *int `json:"slop,omitempty"`

	ZeroTermsQuery *zerotermsquery.ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

func (s *MatchPhraseQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMatchPhraseQuery() *MatchPhraseQuery { _ = "STUB: not implemented"; return nil }

type MatchPhraseQueryVariant interface {
	MatchPhraseQueryCaster() *MatchPhraseQuery
}

func (s *MatchPhraseQuery) MatchPhraseQueryCaster() *MatchPhraseQuery {
	_ = "STUB: not implemented"
	return nil
}
