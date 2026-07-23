package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/remoteclusterprivilege"
)

type _remoteClusterPrivileges struct {
	v *types.RemoteClusterPrivileges
}

func NewRemoteClusterPrivileges() *_remoteClusterPrivileges { _ = "STUB: not implemented"; return nil }

func (s *_remoteClusterPrivileges) Clusters(names ...string) *_remoteClusterPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteClusterPrivileges) Privileges(privileges ...remoteclusterprivilege.RemoteClusterPrivilege) *_remoteClusterPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteClusterPrivileges) RemoteClusterPrivilegesCaster() *types.RemoteClusterPrivileges {
	_ = "STUB: not implemented"
	return nil
}
