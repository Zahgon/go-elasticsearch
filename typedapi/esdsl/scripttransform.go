package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _scriptTransform struct {
	v *types.ScriptTransform
}

func NewScriptTransform() *_scriptTransform { _ = "STUB: not implemented"; return nil }

func (s *_scriptTransform) Id(id string) *_scriptTransform { _ = "STUB: not implemented"; return nil }

func (s *_scriptTransform) Lang(lang string) *_scriptTransform {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptTransform) Params(params map[string]json.RawMessage) *_scriptTransform {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptTransform) AddParam(key string, value json.RawMessage) *_scriptTransform {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptTransform) Source(scriptsource types.ScriptSourceVariant) *_scriptTransform {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptTransform) TransformContainerCaster() *types.TransformContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptTransform) ScriptTransformCaster() *types.ScriptTransform {
	_ = "STUB: not implemented"
	return nil
}
