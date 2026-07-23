package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/remoteclusterprivilege"
)

type RemoteClusterPrivileges struct {
	Clusters []string `json:"clusters"`

	Privileges []remoteclusterprivilege.RemoteClusterPrivilege `json:"privileges"`
}

func (s *RemoteClusterPrivileges) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRemoteClusterPrivileges() *RemoteClusterPrivileges { _ = "STUB: not implemented"; return nil }

type RemoteClusterPrivilegesVariant interface {
	RemoteClusterPrivilegesCaster() *RemoteClusterPrivileges
}

func (s *RemoteClusterPrivileges) RemoteClusterPrivilegesCaster() *RemoteClusterPrivileges {
	_ = "STUB: not implemented"
	return nil
}
