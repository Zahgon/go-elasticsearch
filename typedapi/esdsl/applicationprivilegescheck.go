package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _applicationPrivilegesCheck struct {
	v *types.ApplicationPrivilegesCheck
}

func NewApplicationPrivilegesCheck(application string) *_applicationPrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}

func (s *_applicationPrivilegesCheck) Application(application string) *_applicationPrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}

func (s *_applicationPrivilegesCheck) Privileges(privileges ...string) *_applicationPrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}

func (s *_applicationPrivilegesCheck) Resources(resources ...string) *_applicationPrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}

func (s *_applicationPrivilegesCheck) ApplicationPrivilegesCheckCaster() *types.ApplicationPrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}
