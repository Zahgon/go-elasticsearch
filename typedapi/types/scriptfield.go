package types

type ScriptField struct {
	IgnoreFailure *bool  `json:"ignore_failure,omitempty"`
	Script        Script `json:"script"`
}

func (s *ScriptField) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewScriptField() *ScriptField { _ = "STUB: not implemented"; return nil }

type ScriptFieldVariant interface {
	ScriptFieldCaster() *ScriptField
}

func (s *ScriptField) ScriptFieldCaster() *ScriptField { _ = "STUB: not implemented"; return nil }
