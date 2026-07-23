package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexprivilege"
)

type RemoteIndicesPrivileges struct {
	AllowRestrictedIndices *bool `json:"allow_restricted_indices,omitempty"`

	Clusters []string `json:"clusters"`

	FieldSecurity *FieldSecurity `json:"field_security,omitempty"`

	Names []string `json:"names"`

	Privileges []indexprivilege.IndexPrivilege `json:"privileges"`

	Query IndicesPrivilegesQuery `json:"query,omitempty"`
}

func (s *RemoteIndicesPrivileges) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRemoteIndicesPrivileges() *RemoteIndicesPrivileges { _ = "STUB: not implemented"; return nil }

type RemoteIndicesPrivilegesVariant interface {
	RemoteIndicesPrivilegesCaster() *RemoteIndicesPrivileges
}

func (s *RemoteIndicesPrivileges) RemoteIndicesPrivilegesCaster() *RemoteIndicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}
