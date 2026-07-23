package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
)

type Role struct {
	Applications      []ApplicationPrivileges                   `json:"applications"`
	Cluster           []clusterprivilege.ClusterPrivilege       `json:"cluster"`
	Description       *string                                   `json:"description,omitempty"`
	Global            map[string]map[string]map[string][]string `json:"global,omitempty"`
	Indices           []IndicesPrivilegesRead                   `json:"indices"`
	Metadata          Metadata                                  `json:"metadata"`
	RemoteCluster     []RemoteClusterPrivileges                 `json:"remote_cluster,omitempty"`
	RemoteIndices     []RemoteIndicesPrivileges                 `json:"remote_indices,omitempty"`
	RoleTemplates     []RoleTemplate                            `json:"role_templates,omitempty"`
	RunAs             []string                                  `json:"run_as,omitempty"`
	TransientMetadata map[string]json.RawMessage                `json:"transient_metadata,omitempty"`
}

func (s *Role) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRole() *Role { _ = "STUB: not implemented"; return nil }
