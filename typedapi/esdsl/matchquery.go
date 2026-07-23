package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/zerotermsquery"
)

type _matchQuery struct {
	k string
	v *types.MatchQuery
}

func NewMatchQuery(field string, query string) *_matchQuery { _ = "STUB: not implemented"; return nil }

func (s *_matchQuery) Analyzer(analyzer string) *_matchQuery { _ = "STUB: not implemented"; return nil }

func (s *_matchQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_matchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchQuery) CutoffFrequency(cutofffrequency types.Float64) *_matchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchQuery) Fuzziness(fuzziness types.FuzzinessVariant) *_matchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchQuery) FuzzyRewrite(multitermqueryrewrite string) *_matchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchQuery) FuzzyTranspositions(fuzzytranspositions bool) *_matchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchQuery) Lenient(lenient bool) *_matchQuery { _ = "STUB: not implemented"; return nil }

func (s *_matchQuery) MaxExpansions(maxexpansions int) *_matchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_matchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchQuery) Operator(operator operator.Operator) *_matchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchQuery) PrefixLength(prefixlength int) *_matchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchQuery) Query(query string) *_matchQuery { _ = "STUB: not implemented"; return nil }

func (s *_matchQuery) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *_matchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchQuery) Boost(boost float32) *_matchQuery { _ = "STUB: not implemented"; return nil }

func (s *_matchQuery) QueryName_(queryname_ string) *_matchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_matchQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func NewSingleMatchQuery() *_matchQuery { _ = "STUB: not implemented"; return nil }

func (s *_matchQuery) MatchQueryCaster() *types.MatchQuery { _ = "STUB: not implemented"; return nil }
