package types

type ApplicationGlobalUserPrivileges struct {
	Manage ManageUserPrivileges `json:"manage"`
}

func NewApplicationGlobalUserPrivileges() *ApplicationGlobalUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}

type ApplicationGlobalUserPrivilegesVariant interface {
	ApplicationGlobalUserPrivilegesCaster() *ApplicationGlobalUserPrivileges
}

func (s *ApplicationGlobalUserPrivileges) ApplicationGlobalUserPrivilegesCaster() *ApplicationGlobalUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}
