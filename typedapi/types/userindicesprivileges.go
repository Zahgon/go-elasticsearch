package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexprivilege"
)

type UserIndicesPrivileges struct {
	AllowRestrictedIndices bool `json:"allow_restricted_indices"`

	FieldSecurity []FieldSecurity `json:"field_security,omitempty"`

	Names []string `json:"names"`

	Privileges []indexprivilege.IndexPrivilege `json:"privileges"`

	Query []IndicesPrivilegesQuery `json:"query,omitempty"`
}

func (s *UserIndicesPrivileges) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUserIndicesPrivileges() *UserIndicesPrivileges { _ = "STUB: not implemented"; return nil }
