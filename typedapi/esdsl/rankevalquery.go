package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rankEvalQuery struct {
	v *types.RankEvalQuery
}

func NewRankEvalQuery(query types.QueryVariant) *_rankEvalQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalQuery) Query(query types.QueryVariant) *_rankEvalQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalQuery) Size(size int) *_rankEvalQuery { _ = "STUB: not implemented"; return nil }

func (s *_rankEvalQuery) RankEvalQueryCaster() *types.RankEvalQuery {
	_ = "STUB: not implemented"
	return nil
}
