package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexprivilege"
)

type IndicesPrivileges struct {
	AllowRestrictedIndices *bool `json:"allow_restricted_indices,omitempty"`

	FieldSecurity *FieldSecurity `json:"field_security,omitempty"`

	Names []string `json:"names"`

	Privileges []indexprivilege.IndexPrivilege `json:"privileges"`

	Query IndicesPrivilegesQuery `json:"query,omitempty"`
}

func (s *IndicesPrivileges) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndicesPrivileges() *IndicesPrivileges { _ = "STUB: not implemented"; return nil }

type IndicesPrivilegesVariant interface {
	IndicesPrivilegesCaster() *IndicesPrivileges
}

func (s *IndicesPrivileges) IndicesPrivilegesCaster() *IndicesPrivileges {
	_ = "STUB: not implemented"
	return nil
}
