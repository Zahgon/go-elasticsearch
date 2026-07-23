package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/templateformat"
)

type _roleTemplate struct {
	v *types.RoleTemplate
}

func NewRoleTemplate(template types.ScriptVariant) *_roleTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleTemplate) Format(format templateformat.TemplateFormat) *_roleTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleTemplate) Template(template types.ScriptVariant) *_roleTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleTemplate) RoleTemplateCaster() *types.RoleTemplate {
	_ = "STUB: not implemented"
	return nil
}
