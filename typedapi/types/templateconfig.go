package types

import (
	"encoding/json"
)

type TemplateConfig struct {
	Explain *bool `json:"explain,omitempty"`

	Id *string `json:"id,omitempty"`

	Params map[string]json.RawMessage `json:"params,omitempty"`

	Profile *bool `json:"profile,omitempty"`

	Source ScriptSource `json:"source,omitempty"`
}

func (s *TemplateConfig) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTemplateConfig() *TemplateConfig { _ = "STUB: not implemented"; return nil }

type TemplateConfigVariant interface {
	TemplateConfigCaster() *TemplateConfig
}

func (s *TemplateConfig) TemplateConfigCaster() *TemplateConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *TemplateConfig) RequestItemCaster() *RequestItem { _ = "STUB: not implemented"; return nil }
