package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
)

type QueryRole struct {
	Applications []ApplicationPrivileges `json:"applications,omitempty"`

	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster,omitempty"`

	Description *string `json:"description,omitempty"`

	Global []GlobalPrivilege `json:"global,omitempty"`

	Indices []IndicesPrivileges `json:"indices,omitempty"`

	Metadata Metadata `json:"metadata,omitempty"`

	Name string `json:"name"`

	RemoteCluster []RemoteClusterPrivileges `json:"remote_cluster,omitempty"`

	RemoteIndices []RemoteIndicesPrivileges `json:"remote_indices,omitempty"`

	Restriction *Restriction `json:"restriction,omitempty"`

	RunAs             []string                   `json:"run_as,omitempty"`
	Sort_             []FieldValue               `json:"_sort,omitempty"`
	TransientMetadata map[string]json.RawMessage `json:"transient_metadata,omitempty"`
}

func (s *QueryRole) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewQueryRole() *QueryRole { _ = "STUB: not implemented"; return nil }
