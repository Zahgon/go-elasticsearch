package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/zerotermsquery"
)

type _matchPhraseQuery struct {
	k string
	v *types.MatchPhraseQuery
}

func NewMatchPhraseQuery(field string, query string) *_matchPhraseQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhraseQuery) Analyzer(analyzer string) *_matchPhraseQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhraseQuery) Query(query string) *_matchPhraseQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhraseQuery) Slop(slop int) *_matchPhraseQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhraseQuery) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *_matchPhraseQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhraseQuery) Boost(boost float32) *_matchPhraseQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhraseQuery) QueryName_(queryname_ string) *_matchPhraseQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchPhraseQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func NewSingleMatchPhraseQuery() *_matchPhraseQuery { _ = "STUB: not implemented"; return nil }

func (s *_matchPhraseQuery) MatchPhraseQueryCaster() *types.MatchPhraseQuery {
	_ = "STUB: not implemented"
	return nil
}
