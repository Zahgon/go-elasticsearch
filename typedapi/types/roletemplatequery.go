package types

type RoleTemplateQuery struct {
	Template *RoleTemplateScript `json:"template,omitempty"`
}

func NewRoleTemplateQuery() *RoleTemplateQuery { _ = "STUB: not implemented"; return nil }

type RoleTemplateQueryVariant interface {
	RoleTemplateQueryCaster() *RoleTemplateQuery
}

func (s *RoleTemplateQuery) RoleTemplateQueryCaster() *RoleTemplateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *RoleTemplateQuery) IndicesPrivilegesQueryCaster() *IndicesPrivilegesQuery {
	_ = "STUB: not implemented"
	return nil
}
