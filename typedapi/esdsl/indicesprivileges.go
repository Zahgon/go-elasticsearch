package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexprivilege"
)

type _indicesPrivileges struct {
	v *types.IndicesPrivileges
}

func NewIndicesPrivileges() *_indicesPrivileges { _ = "STUB: not implemented"; return nil }

func (s *_indicesPrivileges) AllowRestrictedIndices(allowrestrictedindices bool) *_indicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesPrivileges) FieldSecurity(fieldsecurity types.FieldSecurityVariant) *_indicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesPrivileges) Names(names ...string) *_indicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesPrivileges) Privileges(privileges ...indexprivilege.IndexPrivilege) *_indicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesPrivileges) Query(indicesprivilegesquery types.IndicesPrivilegesQueryVariant) *_indicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesPrivileges) IndicesPrivilegesCaster() *types.IndicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}
