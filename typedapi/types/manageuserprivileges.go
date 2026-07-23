package types

type ManageUserPrivileges struct {
	Applications []string `json:"applications"`
}

func NewManageUserPrivileges() *ManageUserPrivileges { _ = "STUB: not implemented"; return nil }

type ManageUserPrivilegesVariant interface {
	ManageUserPrivilegesCaster() *ManageUserPrivileges
}

func (s *ManageUserPrivileges) ManageUserPrivilegesCaster() *ManageUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}
