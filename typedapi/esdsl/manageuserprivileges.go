package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _manageUserPrivileges struct {
	v *types.ManageUserPrivileges
}

func NewManageUserPrivileges() *_manageUserPrivileges { _ = "STUB: not implemented"; return nil }

func (s *_manageUserPrivileges) Applications(applications ...string) *_manageUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_manageUserPrivileges) ManageUserPrivilegesCaster() *types.ManageUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}
