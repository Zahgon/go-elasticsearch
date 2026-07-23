package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _functionScore struct {
	v *types.FunctionScore
}

func NewFunctionScore() *_functionScore { _ = "STUB: not implemented"; return nil }

func (s *_functionScore) Exp(decayfunction types.DecayFunctionVariant) *_functionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScore) FieldValueFactor(fieldvaluefactor types.FieldValueFactorScoreFunctionVariant) *_functionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScore) Filter(filter types.QueryVariant) *_functionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScore) Gauss(decayfunction types.DecayFunctionVariant) *_functionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScore) Linear(decayfunction types.DecayFunctionVariant) *_functionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScore) Name_(name_ string) *_functionScore { _ = "STUB: not implemented"; return nil }

func (s *_functionScore) RandomScore(randomscore types.RandomScoreFunctionVariant) *_functionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScore) ScriptScore(scriptscore types.ScriptScoreFunctionVariant) *_functionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScore) Weight(weight types.Float64) *_functionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_functionScore) FunctionScoreCaster() *types.FunctionScore {
	_ = "STUB: not implemented"
	return nil
}
