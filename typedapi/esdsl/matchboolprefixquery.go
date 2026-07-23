package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
)

type _matchBoolPrefixQuery struct {
	k string
	v *types.MatchBoolPrefixQuery
}

func NewMatchBoolPrefixQuery(field string, query string) *_matchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchBoolPrefixQuery) Analyzer(analyzer string) *_matchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchBoolPrefixQuery) Fuzziness(fuzziness types.FuzzinessVariant) *_matchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchBoolPrefixQuery) FuzzyRewrite(multitermqueryrewrite string) *_matchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchBoolPrefixQuery) FuzzyTranspositions(fuzzytranspositions bool) *_matchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchBoolPrefixQuery) MaxExpansions(maxexpansions int) *_matchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchBoolPrefixQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_matchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchBoolPrefixQuery) Operator(operator operator.Operator) *_matchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchBoolPrefixQuery) PrefixLength(prefixlength int) *_matchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchBoolPrefixQuery) Query(query string) *_matchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchBoolPrefixQuery) Boost(boost float32) *_matchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchBoolPrefixQuery) QueryName_(queryname_ string) *_matchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchBoolPrefixQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func NewSingleMatchBoolPrefixQuery() *_matchBoolPrefixQuery { _ = "STUB: not implemented"; return nil }

func (s *_matchBoolPrefixQuery) MatchBoolPrefixQueryCaster() *types.MatchBoolPrefixQuery {
	_ = "STUB: not implemented"
	return nil
}
