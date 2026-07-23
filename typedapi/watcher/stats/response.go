package stats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ClusterName     string                   `json:"cluster_name"`
	ManuallyStopped bool                     `json:"manually_stopped"`
	NodeStats       types.NodeStatistics     `json:"_nodes"`
	Stats           []types.WatcherNodeStats `json:"stats"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
