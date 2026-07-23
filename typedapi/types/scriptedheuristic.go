package types

type ScriptedHeuristic struct {
	Script Script `json:"script"`
}

func NewScriptedHeuristic() *ScriptedHeuristic { _ = "STUB: not implemented"; return nil }

type ScriptedHeuristicVariant interface {
	ScriptedHeuristicCaster() *ScriptedHeuristic
}

func (s *ScriptedHeuristic) ScriptedHeuristicCaster() *ScriptedHeuristic {
	_ = "STUB: not implemented"
	return nil
}
