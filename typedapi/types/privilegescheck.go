package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
)

type PrivilegesCheck struct {
	Application []ApplicationPrivilegesCheck `json:"application,omitempty"`

	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster,omitempty"`
	Index   []IndexPrivilegesCheck              `json:"index,omitempty"`
}

func NewPrivilegesCheck() *PrivilegesCheck { _ = "STUB: not implemented"; return nil }

type PrivilegesCheckVariant interface {
	PrivilegesCheckCaster() *PrivilegesCheck
}

func (s *PrivilegesCheck) PrivilegesCheckCaster() *PrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}
