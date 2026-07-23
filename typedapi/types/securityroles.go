package types

type SecurityRoles struct {
	Dls    SecurityRolesDls    `json:"dls"`
	File   SecurityRolesFile   `json:"file"`
	Native SecurityRolesNative `json:"native"`
}

func NewSecurityRoles() *SecurityRoles { _ = "STUB: not implemented"; return nil }
