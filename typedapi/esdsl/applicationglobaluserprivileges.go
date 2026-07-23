package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _applicationGlobalUserPrivileges struct {
	v *types.ApplicationGlobalUserPrivileges
}

func NewApplicationGlobalUserPrivileges(manage types.ManageUserPrivilegesVariant) *_applicationGlobalUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_applicationGlobalUserPrivileges) Manage(manage types.ManageUserPrivilegesVariant) *_applicationGlobalUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_applicationGlobalUserPrivileges) ApplicationGlobalUserPrivilegesCaster() *types.ApplicationGlobalUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}
