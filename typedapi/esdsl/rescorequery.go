package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scoremode"
)

type _rescoreQuery struct {
	v *types.RescoreQuery
}

func NewRescoreQuery(query types.QueryVariant) *_rescoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescoreQuery) Query(query types.QueryVariant) *_rescoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescoreQuery) QueryWeight(queryweight types.Float64) *_rescoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescoreQuery) RescoreQueryWeight(rescorequeryweight types.Float64) *_rescoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescoreQuery) ScoreMode(scoremode scoremode.ScoreMode) *_rescoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescoreQuery) RescoreCaster() *types.Rescore { _ = "STUB: not implemented"; return nil }

func (s *_rescoreQuery) RescoreQueryCaster() *types.RescoreQuery {
	_ = "STUB: not implemented"
	return nil
}
