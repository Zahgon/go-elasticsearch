package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type ScriptCondition struct {
	Id     *string                        `json:"id,omitempty"`
	Lang   *scriptlanguage.ScriptLanguage `json:"lang,omitempty"`
	Params map[string]json.RawMessage     `json:"params,omitempty"`
	Source ScriptSource                   `json:"source,omitempty"`
}

func (s *ScriptCondition) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewScriptCondition() *ScriptCondition { _ = "STUB: not implemented"; return nil }

type ScriptConditionVariant interface {
	ScriptConditionCaster() *ScriptCondition
}

func (s *ScriptCondition) ScriptConditionCaster() *ScriptCondition {
	_ = "STUB: not implemented"
	return nil
}
