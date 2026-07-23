package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/zerotermsquery"
)

type _matchPhrasePrefixQuery struct {
	k string
	v *types.MatchPhrasePrefixQuery
}

func NewMatchPhrasePrefixQuery(field string, query string) *_matchPhrasePrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhrasePrefixQuery) Analyzer(analyzer string) *_matchPhrasePrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhrasePrefixQuery) MaxExpansions(maxexpansions int) *_matchPhrasePrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhrasePrefixQuery) Query(query string) *_matchPhrasePrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhrasePrefixQuery) Slop(slop int) *_matchPhrasePrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhrasePrefixQuery) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *_matchPhrasePrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhrasePrefixQuery) Boost(boost float32) *_matchPhrasePrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhrasePrefixQuery) QueryName_(queryname_ string) *_matchPhrasePrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhrasePrefixQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func NewSingleMatchPhrasePrefixQuery() *_matchPhrasePrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhrasePrefixQuery) MatchPhrasePrefixQueryCaster() *types.MatchPhrasePrefixQuery {
	_ = "STUB: not implemented"
	return nil
}
