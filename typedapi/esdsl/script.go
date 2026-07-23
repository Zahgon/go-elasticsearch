package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type _script struct {
	v *types.Script
}

func NewScript() *_script { _ = "STUB: not implemented"; return nil }

func (s *_script) Id(id string) *_script { _ = "STUB: not implemented"; return nil }

func (s *_script) Lang(lang scriptlanguage.ScriptLanguage) *_script {
	_ = "STUB: not implemented"
	return nil
}

func (s *_script) Options(options map[string]string) *_script {
	_ = "STUB: not implemented"
	return nil
}

func (s *_script) AddOption(key string, value string) *_script {
	_ = "STUB: not implemented"
	return nil
}

func (s *_script) Params(params map[string]json.RawMessage) *_script {
	_ = "STUB: not implemented"
	return nil
}

func (s *_script) AddParam(key string, value json.RawMessage) *_script {
	_ = "STUB: not implemented"
	return nil
}

func (s *_script) Source(scriptsource types.ScriptSourceVariant) *_script {
	_ = "STUB: not implemented"
	return nil
}

func (s *_script) MultiTermLookupCaster() *types.MultiTermLookup {
	_ = "STUB: not implemented"
	return nil
}

func (s *_script) IntervalsFilterCaster() *types.IntervalsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_script) ScriptCaster() *types.Script { _ = "STUB: not implemented"; return nil }
