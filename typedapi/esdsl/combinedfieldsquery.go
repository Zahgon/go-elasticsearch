package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/combinedfieldsoperator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/combinedfieldszeroterms"
)

type _combinedFieldsQuery struct {
	v *types.CombinedFieldsQuery
}

func NewCombinedFieldsQuery(query string) *_combinedFieldsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_combinedFieldsQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_combinedFieldsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_combinedFieldsQuery) Fields(fields ...string) *_combinedFieldsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_combinedFieldsQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_combinedFieldsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_combinedFieldsQuery) Operator(operator combinedfieldsoperator.CombinedFieldsOperator) *_combinedFieldsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_combinedFieldsQuery) Query(query string) *_combinedFieldsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_combinedFieldsQuery) ZeroTermsQuery(zerotermsquery combinedfieldszeroterms.CombinedFieldsZeroTerms) *_combinedFieldsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_combinedFieldsQuery) Boost(boost float32) *_combinedFieldsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_combinedFieldsQuery) QueryName_(queryname_ string) *_combinedFieldsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_combinedFieldsQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_combinedFieldsQuery) CombinedFieldsQueryCaster() *types.CombinedFieldsQuery {
	_ = "STUB: not implemented"
	return nil
}
