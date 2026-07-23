package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexprivilege"
)

type IndexPrivilegesCheck struct {
	AllowRestrictedIndices *bool `json:"allow_restricted_indices,omitempty"`

	Names []string `json:"names"`

	Privileges []indexprivilege.IndexPrivilege `json:"privileges"`
}

func (s *IndexPrivilegesCheck) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexPrivilegesCheck() *IndexPrivilegesCheck { _ = "STUB: not implemented"; return nil }

type IndexPrivilegesCheckVariant interface {
	IndexPrivilegesCheckCaster() *IndexPrivilegesCheck
}

func (s *IndexPrivilegesCheck) IndexPrivilegesCheckCaster() *IndexPrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}
