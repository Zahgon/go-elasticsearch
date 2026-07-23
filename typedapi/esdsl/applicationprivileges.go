package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _applicationPrivileges struct {
	v *types.ApplicationPrivileges
}

func NewApplicationPrivileges(application string) *_applicationPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_applicationPrivileges) Application(application string) *_applicationPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_applicationPrivileges) Privileges(privileges ...string) *_applicationPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_applicationPrivileges) Resources(resources ...string) *_applicationPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_applicationPrivileges) ApplicationPrivilegesCaster() *types.ApplicationPrivileges {
	_ = "STUB: not implemented"
	return nil
}
