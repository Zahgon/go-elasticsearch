package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/textquerytype"
)

type _queryStringQuery struct {
	v *types.QueryStringQuery
}

func NewQueryStringQuery(query string) *_queryStringQuery { _ = "STUB: not implemented"; return nil }

func (s *_queryStringQuery) AllowLeadingWildcard(allowleadingwildcard bool) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) AnalyzeWildcard(analyzewildcard bool) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) Analyzer(analyzer string) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) DefaultField(field string) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) DefaultOperator(defaultoperator operator.Operator) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) EnablePositionIncrements(enablepositionincrements bool) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) Escape(escape bool) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) Fields(fields ...string) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) Fuzziness(fuzziness types.FuzzinessVariant) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) FuzzyMaxExpansions(fuzzymaxexpansions int) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) FuzzyPrefixLength(fuzzyprefixlength int) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) FuzzyRewrite(multitermqueryrewrite string) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) FuzzyTranspositions(fuzzytranspositions bool) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) Lenient(lenient bool) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) MaxDeterminizedStates(maxdeterminizedstates int) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) PhraseSlop(phraseslop types.Float64) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) Query(query string) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) QuoteAnalyzer(quoteanalyzer string) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) QuoteFieldSuffix(quotefieldsuffix string) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) Rewrite(multitermqueryrewrite string) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) TieBreaker(tiebreaker types.Float64) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) TimeZone(timezone string) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) Type(type_ textquerytype.TextQueryType) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) Boost(boost float32) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) QueryName_(queryname_ string) *_queryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryStringQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_queryStringQuery) QueryStringQueryCaster() *types.QueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}
