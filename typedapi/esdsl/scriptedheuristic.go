package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scriptedHeuristic struct {
	v *types.ScriptedHeuristic
}

func NewScriptedHeuristic(script types.ScriptVariant) *_scriptedHeuristic {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptedHeuristic) Script(script types.ScriptVariant) *_scriptedHeuristic {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptedHeuristic) ScriptedHeuristicCaster() *types.ScriptedHeuristic {
	_ = "STUB: not implemented"
	return nil
}
