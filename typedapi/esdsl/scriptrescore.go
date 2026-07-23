package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scriptRescore struct {
	v *types.ScriptRescore
}

func NewScriptRescore(script types.ScriptVariant) *_scriptRescore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptRescore) Script(script types.ScriptVariant) *_scriptRescore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptRescore) RescoreCaster() *types.Rescore { _ = "STUB: not implemented"; return nil }

func (s *_scriptRescore) ScriptRescoreCaster() *types.ScriptRescore {
	_ = "STUB: not implemented"
	return nil
}
