package putrole

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
)

type Request struct {
	Applications []types.ApplicationPrivileges `json:"applications,omitempty"`

	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster,omitempty"`

	Description *string `json:"description,omitempty"`

	Global map[string]json.RawMessage `json:"global,omitempty"`

	Indices []types.IndicesPrivileges `json:"indices,omitempty"`

	Metadata types.Metadata `json:"metadata,omitempty"`

	RemoteCluster []types.RemoteClusterPrivileges `json:"remote_cluster,omitempty"`

	RemoteIndices []types.RemoteIndicesPrivileges `json:"remote_indices,omitempty"`

	RunAs []string `json:"run_as,omitempty"`

	TransientMetadata map[string]json.RawMessage `json:"transient_metadata,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
