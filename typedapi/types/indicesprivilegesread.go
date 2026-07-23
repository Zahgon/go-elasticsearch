package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexprivilege"
)

type IndicesPrivilegesRead struct {
	AllowRestrictedIndices *bool `json:"allow_restricted_indices,omitempty"`

	FieldSecurity *FieldSecurity `json:"field_security,omitempty"`

	ImplicitlyGranted *bool `json:"implicitly_granted,omitempty"`

	Names []string `json:"names"`

	Privileges []indexprivilege.IndexPrivilege `json:"privileges"`

	Query IndicesPrivilegesQuery `json:"query,omitempty"`
}

func (s *IndicesPrivilegesRead) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndicesPrivilegesRead() *IndicesPrivilegesRead { _ = "STUB: not implemented"; return nil }
