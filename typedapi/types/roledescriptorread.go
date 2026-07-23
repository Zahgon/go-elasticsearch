package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
)

type RoleDescriptorRead struct {
	Applications []ApplicationPrivileges `json:"applications,omitempty"`

	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster"`

	Description *string `json:"description,omitempty"`

	Global []GlobalPrivilege `json:"global,omitempty"`

	Indices []IndicesPrivileges `json:"indices"`

	Metadata Metadata `json:"metadata,omitempty"`

	RemoteCluster []RemoteClusterPrivileges `json:"remote_cluster,omitempty"`

	RemoteIndices []RemoteIndicesPrivileges `json:"remote_indices,omitempty"`

	Restriction *Restriction `json:"restriction,omitempty"`

	RunAs             []string                   `json:"run_as,omitempty"`
	TransientMetadata map[string]json.RawMessage `json:"transient_metadata,omitempty"`
}

func (s *RoleDescriptorRead) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRoleDescriptorRead() *RoleDescriptorRead { _ = "STUB: not implemented"; return nil }
