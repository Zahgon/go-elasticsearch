package types

type ScriptScoreFunction struct {
	Script Script `json:"script"`
}

func NewScriptScoreFunction() *ScriptScoreFunction { _ = "STUB: not implemented"; return nil }

type ScriptScoreFunctionVariant interface {
	ScriptScoreFunctionCaster() *ScriptScoreFunction
}

func (s *ScriptScoreFunction) ScriptScoreFunctionCaster() *ScriptScoreFunction {
	_ = "STUB: not implemented"
	return nil
}
