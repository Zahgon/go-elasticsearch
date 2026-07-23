package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexprivilege"
)

type _remoteIndicesPrivileges struct {
	v *types.RemoteIndicesPrivileges
}

func NewRemoteIndicesPrivileges() *_remoteIndicesPrivileges { _ = "STUB: not implemented"; return nil }

func (s *_remoteIndicesPrivileges) AllowRestrictedIndices(allowrestrictedindices bool) *_remoteIndicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteIndicesPrivileges) Clusters(names ...string) *_remoteIndicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteIndicesPrivileges) FieldSecurity(fieldsecurity types.FieldSecurityVariant) *_remoteIndicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteIndicesPrivileges) Names(names ...string) *_remoteIndicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteIndicesPrivileges) Privileges(privileges ...indexprivilege.IndexPrivilege) *_remoteIndicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteIndicesPrivileges) Query(indicesprivilegesquery types.IndicesPrivilegesQueryVariant) *_remoteIndicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteIndicesPrivileges) RemoteIndicesPrivilegesCaster() *types.RemoteIndicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}
