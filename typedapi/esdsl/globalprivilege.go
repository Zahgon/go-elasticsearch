package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _globalPrivilege struct {
	v *types.GlobalPrivilege
}

func NewGlobalPrivilege(application types.ApplicationGlobalUserPrivilegesVariant) *_globalPrivilege {
	_ = "STUB: not implemented"
	return nil
}

func (s *_globalPrivilege) Application(application types.ApplicationGlobalUserPrivilegesVariant) *_globalPrivilege {
	_ = "STUB: not implemented"
	return nil
}

func (s *_globalPrivilege) GlobalPrivilegeCaster() *types.GlobalPrivilege {
	_ = "STUB: not implemented"
	return nil
}
