package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
)

type _simpleQueryStringQuery struct {
	v *types.SimpleQueryStringQuery
}

func NewSimpleQueryStringQuery(query string) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) AnalyzeWildcard(analyzewildcard bool) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) Analyzer(analyzer string) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) DefaultOperator(defaultoperator operator.Operator) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) Fields(fields ...string) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) Flags(simplequerystringflags types.PipeSeparatedFlagsSimpleQueryStringFlag) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) FuzzyMaxExpansions(fuzzymaxexpansions int) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) FuzzyPrefixLength(fuzzyprefixlength int) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) FuzzyTranspositions(fuzzytranspositions bool) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) Lenient(lenient bool) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) Query(query string) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) QuoteFieldSuffix(quotefieldsuffix string) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) Boost(boost float32) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) QueryName_(queryname_ string) *_simpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_simpleQueryStringQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleQueryStringQuery) SimpleQueryStringQueryCaster() *types.SimpleQueryStringQuery {
	_ = "STUB: not implemented"
	return nil
}
