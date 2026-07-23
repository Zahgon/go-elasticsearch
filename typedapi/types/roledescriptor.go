package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
)

type RoleDescriptor struct {
	Applications []ApplicationPrivileges `json:"applications,omitempty"`

	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster,omitempty"`

	Description *string `json:"description,omitempty"`

	Global []GlobalPrivilege `json:"global,omitempty"`

	Indices []IndicesPrivileges `json:"indices,omitempty"`

	Metadata Metadata `json:"metadata,omitempty"`

	RemoteCluster []RemoteClusterPrivileges `json:"remote_cluster,omitempty"`

	RemoteIndices []RemoteIndicesPrivileges `json:"remote_indices,omitempty"`

	Restriction *Restriction `json:"restriction,omitempty"`

	RunAs             []string                   `json:"run_as,omitempty"`
	TransientMetadata map[string]json.RawMessage `json:"transient_metadata,omitempty"`
}

func (s *RoleDescriptor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRoleDescriptor() *RoleDescriptor { _ = "STUB: not implemented"; return nil }

type RoleDescriptorVariant interface {
	RoleDescriptorCaster() *RoleDescriptor
}

func (s *RoleDescriptor) RoleDescriptorCaster() *RoleDescriptor {
	_ = "STUB: not implemented"
	return nil
}
