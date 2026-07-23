package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type _scriptCondition struct {
	v *types.ScriptCondition
}

func NewScriptCondition() *_scriptCondition { _ = "STUB: not implemented"; return nil }

func (s *_scriptCondition) Id(id string) *_scriptCondition { _ = "STUB: not implemented"; return nil }

func (s *_scriptCondition) Lang(lang scriptlanguage.ScriptLanguage) *_scriptCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptCondition) Params(params map[string]json.RawMessage) *_scriptCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptCondition) AddParam(key string, value json.RawMessage) *_scriptCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptCondition) Source(scriptsource types.ScriptSourceVariant) *_scriptCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptCondition) WatcherConditionCaster() *types.WatcherCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptCondition) ScriptConditionCaster() *types.ScriptCondition {
	_ = "STUB: not implemented"
	return nil
}
