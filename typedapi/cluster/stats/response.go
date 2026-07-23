package stats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/healthstatus"
)

type Response struct {
	Ccs types.CCSStats `json:"ccs"`

	ClusterName string `json:"cluster_name"`

	ClusterUuid string `json:"cluster_uuid"`

	Indices types.ClusterIndices `json:"indices"`

	NodeStats *types.NodeStatistics `json:"_nodes,omitempty"`

	Nodes types.ClusterNodes `json:"nodes"`

	Repositories map[string]map[string]int64 `json:"repositories"`

	Snapshots types.ClusterSnapshotStats `json:"snapshots"`

	Status *healthstatus.HealthStatus `json:"status,omitempty"`

	Timestamp int64 `json:"timestamp"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
