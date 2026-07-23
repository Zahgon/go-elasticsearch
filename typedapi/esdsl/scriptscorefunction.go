package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scriptScoreFunction struct {
	v *types.ScriptScoreFunction
}

func NewScriptScoreFunction(script types.ScriptVariant) *_scriptScoreFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptScoreFunction) Script(script types.ScriptVariant) *_scriptScoreFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptScoreFunction) FunctionScoreCaster() *types.FunctionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptScoreFunction) ScriptScoreFunctionCaster() *types.ScriptScoreFunction {
	_ = "STUB: not implemented"
	return nil
}
