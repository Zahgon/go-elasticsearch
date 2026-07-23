package types

type ScriptRescore struct {
	Script Script `json:"script"`
}

func NewScriptRescore() *ScriptRescore { _ = "STUB: not implemented"; return nil }

type ScriptRescoreVariant interface {
	ScriptRescoreCaster() *ScriptRescore
}

func (s *ScriptRescore) ScriptRescoreCaster() *ScriptRescore { _ = "STUB: not implemented"; return nil }
