package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexprivilege"
)

type _indexPrivilegesCheck struct {
	v *types.IndexPrivilegesCheck
}

func NewIndexPrivilegesCheck() *_indexPrivilegesCheck { _ = "STUB: not implemented"; return nil }

func (s *_indexPrivilegesCheck) AllowRestrictedIndices(allowrestrictedindices bool) *_indexPrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexPrivilegesCheck) Names(indices ...string) *_indexPrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexPrivilegesCheck) Privileges(privileges ...indexprivilege.IndexPrivilege) *_indexPrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexPrivilegesCheck) IndexPrivilegesCheckCaster() *types.IndexPrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}
