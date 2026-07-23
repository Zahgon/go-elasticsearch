package types

type SecurityRoleMapping struct {
	Enabled       bool            `json:"enabled"`
	Metadata      Metadata        `json:"metadata"`
	RoleTemplates []RoleTemplate  `json:"role_templates,omitempty"`
	Roles         []string        `json:"roles,omitempty"`
	Rules         RoleMappingRule `json:"rules"`
}

func (s *SecurityRoleMapping) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSecurityRoleMapping() *SecurityRoleMapping { _ = "STUB: not implemented"; return nil }
