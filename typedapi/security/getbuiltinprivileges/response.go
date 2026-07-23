package getbuiltinprivileges

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/remoteclusterprivilege"
)

type Response struct {
	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster"`

	Index []string `json:"index"`

	RemoteCluster []remoteclusterprivilege.RemoteClusterPrivilege `json:"remote_cluster"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
