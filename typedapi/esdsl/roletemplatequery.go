package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _roleTemplateQuery struct {
	v *types.RoleTemplateQuery
}

func NewRoleTemplateQuery() *_roleTemplateQuery { _ = "STUB: not implemented"; return nil }

func (s *_roleTemplateQuery) Template(template types.RoleTemplateScriptVariant) *_roleTemplateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleTemplateQuery) RoleTemplateQueryCaster() *types.RoleTemplateQuery {
	_ = "STUB: not implemented"
	return nil
}
