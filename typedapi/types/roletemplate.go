package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/templateformat"
)

type RoleTemplate struct {
	Format   *templateformat.TemplateFormat `json:"format,omitempty"`
	Template Script                         `json:"template"`
}

func NewRoleTemplate() *RoleTemplate { _ = "STUB: not implemented"; return nil }

type RoleTemplateVariant interface {
	RoleTemplateCaster() *RoleTemplate
}

func (s *RoleTemplate) RoleTemplateCaster() *RoleTemplate { _ = "STUB: not implemented"; return nil }
