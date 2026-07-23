package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scriptField struct {
	v *types.ScriptField
}

func NewScriptField(script types.ScriptVariant) *_scriptField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptField) IgnoreFailure(ignorefailure bool) *_scriptField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptField) Script(script types.ScriptVariant) *_scriptField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptField) ScriptFieldCaster() *types.ScriptField {
	_ = "STUB: not implemented"
	return nil
}
