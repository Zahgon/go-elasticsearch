package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/zerotermsquery"
)

type MatchPhrasePrefixQuery struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Boost *float32 `json:"boost,omitempty"`

	MaxExpansions *int `json:"max_expansions,omitempty"`

	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`

	Slop *int `json:"slop,omitempty"`

	ZeroTermsQuery *zerotermsquery.ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

func (s *MatchPhrasePrefixQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMatchPhrasePrefixQuery() *MatchPhrasePrefixQuery { _ = "STUB: not implemented"; return nil }

type MatchPhrasePrefixQueryVariant interface {
	MatchPhrasePrefixQueryCaster() *MatchPhrasePrefixQuery
}

func (s *MatchPhrasePrefixQuery) MatchPhrasePrefixQueryCaster() *MatchPhrasePrefixQuery {
	_ = "STUB: not implemented"
	return nil
}
