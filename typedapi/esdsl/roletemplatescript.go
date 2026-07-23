package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type _roleTemplateScript struct {
	v *types.RoleTemplateScript
}

func NewRoleTemplateScript() *_roleTemplateScript { _ = "STUB: not implemented"; return nil }

func (s *_roleTemplateScript) Id(id string) *_roleTemplateScript {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleTemplateScript) Lang(lang scriptlanguage.ScriptLanguage) *_roleTemplateScript {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleTemplateScript) Options(options map[string]string) *_roleTemplateScript {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleTemplateScript) AddOption(key string, value string) *_roleTemplateScript {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleTemplateScript) Params(params map[string]json.RawMessage) *_roleTemplateScript {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleTemplateScript) AddParam(key string, value json.RawMessage) *_roleTemplateScript {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleTemplateScript) Source(roletemplateinlinequery types.RoleTemplateInlineQueryVariant) *_roleTemplateScript {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleTemplateScript) RoleTemplateScriptCaster() *types.RoleTemplateScript {
	_ = "STUB: not implemented"
	return nil
}
