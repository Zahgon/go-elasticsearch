package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type RoleTemplateScript struct {
	Id *string `json:"id,omitempty"`

	Lang    *scriptlanguage.ScriptLanguage `json:"lang,omitempty"`
	Options map[string]string              `json:"options,omitempty"`

	Params map[string]json.RawMessage `json:"params,omitempty"`
	Source RoleTemplateInlineQuery    `json:"source,omitempty"`
}

func (s *RoleTemplateScript) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRoleTemplateScript() *RoleTemplateScript { _ = "STUB: not implemented"; return nil }

type RoleTemplateScriptVariant interface {
	RoleTemplateScriptCaster() *RoleTemplateScript
}

func (s *RoleTemplateScript) RoleTemplateScriptCaster() *RoleTemplateScript {
	_ = "STUB: not implemented"
	return nil
}
