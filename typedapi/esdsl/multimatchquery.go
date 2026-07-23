package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/textquerytype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/zerotermsquery"
)

type _multiMatchQuery struct {
	v *types.MultiMatchQuery
}

func NewMultiMatchQuery(query string) *_multiMatchQuery { _ = "STUB: not implemented"; return nil }

func (s *_multiMatchQuery) Analyzer(analyzer string) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) CutoffFrequency(cutofffrequency types.Float64) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) Fields(fields ...string) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) Fuzziness(fuzziness types.FuzzinessVariant) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) FuzzyRewrite(multitermqueryrewrite string) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) FuzzyTranspositions(fuzzytranspositions bool) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) Lenient(lenient bool) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) MaxExpansions(maxexpansions int) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) Operator(operator operator.Operator) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) PrefixLength(prefixlength int) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) Query(query string) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) Slop(slop int) *_multiMatchQuery { _ = "STUB: not implemented"; return nil }

func (s *_multiMatchQuery) TieBreaker(tiebreaker types.Float64) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) Type(type_ textquerytype.TextQueryType) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) Boost(boost float32) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) QueryName_(queryname_ string) *_multiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multiMatchQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_multiMatchQuery) MultiMatchQueryCaster() *types.MultiMatchQuery {
	_ = "STUB: not implemented"
	return nil
}
