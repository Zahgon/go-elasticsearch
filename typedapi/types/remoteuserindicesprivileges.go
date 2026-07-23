package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexprivilege"
)

type RemoteUserIndicesPrivileges struct {
	AllowRestrictedIndices bool     `json:"allow_restricted_indices"`
	Clusters               []string `json:"clusters"`

	FieldSecurity []FieldSecurity `json:"field_security,omitempty"`

	Names []string `json:"names"`

	Privileges []indexprivilege.IndexPrivilege `json:"privileges"`

	Query []IndicesPrivilegesQuery `json:"query,omitempty"`
}

func (s *RemoteUserIndicesPrivileges) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRemoteUserIndicesPrivileges() *RemoteUserIndicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}
