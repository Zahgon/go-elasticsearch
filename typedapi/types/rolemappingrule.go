package types

type RoleMappingRule struct {
	All    []RoleMappingRule       `json:"all,omitempty"`
	Any    []RoleMappingRule       `json:"any,omitempty"`
	Except *RoleMappingRule        `json:"except,omitempty"`
	Field  map[string][]FieldValue `json:"field,omitempty"`
}

func (s *RoleMappingRule) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRoleMappingRule() *RoleMappingRule { _ = "STUB: not implemented"; return nil }

type RoleMappingRuleVariant interface {
	RoleMappingRuleCaster() *RoleMappingRule
}

func (s *RoleMappingRule) RoleMappingRuleCaster() *RoleMappingRule {
	_ = "STUB: not implemented"
	return nil
}
