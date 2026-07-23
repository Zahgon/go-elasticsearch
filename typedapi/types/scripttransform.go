package types

import (
	"encoding/json"
)

type ScriptTransform struct {
	Id     *string                    `json:"id,omitempty"`
	Lang   *string                    `json:"lang,omitempty"`
	Params map[string]json.RawMessage `json:"params,omitempty"`
	Source ScriptSource               `json:"source,omitempty"`
}

func (s *ScriptTransform) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewScriptTransform() *ScriptTransform { _ = "STUB: not implemented"; return nil }

type ScriptTransformVariant interface {
	ScriptTransformCaster() *ScriptTransform
}

func (s *ScriptTransform) ScriptTransformCaster() *ScriptTransform {
	_ = "STUB: not implemented"
	return nil
}
