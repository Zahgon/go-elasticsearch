package getuserprivileges

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Applications  []types.ApplicationPrivileges       `json:"applications"`
	Cluster       []string                            `json:"cluster"`
	Global        []types.GlobalPrivilege             `json:"global"`
	Indices       []types.UserIndicesPrivileges       `json:"indices"`
	RemoteCluster []types.RemoteClusterPrivileges     `json:"remote_cluster,omitempty"`
	RemoteIndices []types.RemoteUserIndicesPrivileges `json:"remote_indices,omitempty"`
	RunAs         []string                            `json:"run_as"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
