package types

type GlobalPrivilege struct {
	Application ApplicationGlobalUserPrivileges `json:"application"`
}

func NewGlobalPrivilege() *GlobalPrivilege { _ = "STUB: not implemented"; return nil }

type GlobalPrivilegeVariant interface {
	GlobalPrivilegeCaster() *GlobalPrivilege
}

func (s *GlobalPrivilege) GlobalPrivilegeCaster() *GlobalPrivilege {
	_ = "STUB: not implemented"
	return nil
}
