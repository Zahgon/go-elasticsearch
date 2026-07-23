package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _templateConfig struct {
	v *types.TemplateConfig
}

func NewTemplateConfig() *_templateConfig { _ = "STUB: not implemented"; return nil }

func (s *_templateConfig) Explain(explain bool) *_templateConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_templateConfig) Id(id string) *_templateConfig { _ = "STUB: not implemented"; return nil }

func (s *_templateConfig) Params(params map[string]json.RawMessage) *_templateConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_templateConfig) AddParam(key string, value json.RawMessage) *_templateConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_templateConfig) Profile(profile bool) *_templateConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_templateConfig) Source(scriptsource types.ScriptSourceVariant) *_templateConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_templateConfig) TemplateConfigCaster() *types.TemplateConfig {
	_ = "STUB: not implemented"
	return nil
}
