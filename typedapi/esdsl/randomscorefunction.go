package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _randomScoreFunction struct {
	v *types.RandomScoreFunction
}

func NewRandomScoreFunction() *_randomScoreFunction { _ = "STUB: not implemented"; return nil }

func (s *_randomScoreFunction) Field(field string) *_randomScoreFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_randomScoreFunction) Seed(seed string) *_randomScoreFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_randomScoreFunction) FunctionScoreCaster() *types.FunctionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_randomScoreFunction) RandomScoreFunctionCaster() *types.RandomScoreFunction {
	_ = "STUB: not implemented"
	return nil
}
