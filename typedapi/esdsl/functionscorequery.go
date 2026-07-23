package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/functionboostmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/functionscoremode"
)

type _functionScoreQuery struct {
	v *types.FunctionScoreQuery
}

func NewFunctionScoreQuery() *_functionScoreQuery { _ = "STUB: not implemented"; return nil }

func (s *_functionScoreQuery) BoostMode(boostmode functionboostmode.FunctionBoostMode) *_functionScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScoreQuery) Functions(functions ...types.FunctionScoreVariant) *_functionScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScoreQuery) FunctionsValues(functionsvalues []types.FunctionScore) *_functionScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScoreQuery) MaxBoost(maxboost types.Float64) *_functionScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScoreQuery) MinScore(minscore types.Float64) *_functionScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScoreQuery) Query(query types.QueryVariant) *_functionScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScoreQuery) ScoreMode(scoremode functionscoremode.FunctionScoreMode) *_functionScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScoreQuery) Boost(boost float32) *_functionScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScoreQuery) QueryName_(queryname_ string) *_functionScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScoreQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_functionScoreQuery) FunctionScoreQueryCaster() *types.FunctionScoreQuery {
	_ = "STUB: not implemented"
	return nil
}
