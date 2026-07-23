package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
)

type _privilegesCheck struct {
	v *types.PrivilegesCheck
}

func NewPrivilegesCheck() *_privilegesCheck { _ = "STUB: not implemented"; return nil }

func (s *_privilegesCheck) Application(applications ...types.ApplicationPrivilegesCheckVariant) *_privilegesCheck {
	_ = "STUB: not implemented"
	return nil
}

func (s *_privilegesCheck) ApplicationValues(applicationvalues []types.ApplicationPrivilegesCheck) *_privilegesCheck {
	_ = "STUB: not implemented"
	return nil
}

func (s *_privilegesCheck) Cluster(clusters ...clusterprivilege.ClusterPrivilege) *_privilegesCheck {
	_ = "STUB: not implemented"
	return nil
}

func (s *_privilegesCheck) Index(indices ...types.IndexPrivilegesCheckVariant) *_privilegesCheck {
	_ = "STUB: not implemented"
	return nil
}

func (s *_privilegesCheck) IndexValues(indexvalues []types.IndexPrivilegesCheck) *_privilegesCheck {
	_ = "STUB: not implemented"
	return nil
}

func (s *_privilegesCheck) PrivilegesCheckCaster() *types.PrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}
