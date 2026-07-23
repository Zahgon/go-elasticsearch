package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/noderole"
)

type NodeInfo struct {
	Aggregations map[string]NodeInfoAggregation `json:"aggregations,omitempty"`
	Attributes   map[string]string              `json:"attributes"`
	BuildFlavor  string                         `json:"build_flavor"`

	BuildHash         string         `json:"build_hash"`
	BuildType         string         `json:"build_type"`
	ComponentVersions map[string]int `json:"component_versions"`

	Host         string          `json:"host"`
	Http         *NodeInfoHttp   `json:"http,omitempty"`
	IndexVersion int64           `json:"index_version"`
	Ingest       *NodeInfoIngest `json:"ingest,omitempty"`

	Ip      string        `json:"ip"`
	Jvm     *NodeJvmInfo  `json:"jvm,omitempty"`
	Modules []PluginStats `json:"modules,omitempty"`

	Name                string                        `json:"name"`
	Os                  *NodeOperatingSystemInfo      `json:"os,omitempty"`
	Plugins             []PluginStats                 `json:"plugins,omitempty"`
	Process             *NodeProcessInfo              `json:"process,omitempty"`
	RemoteClusterServer *RemoveClusterServer          `json:"remote_cluster_server,omitempty"`
	Roles               []noderole.NodeRole           `json:"roles"`
	Settings            *NodeInfoSettings             `json:"settings,omitempty"`
	ThreadPool          map[string]NodeThreadPoolInfo `json:"thread_pool,omitempty"`

	TotalIndexingBuffer *int64 `json:"total_indexing_buffer,omitempty"`

	TotalIndexingBufferInBytes ByteSize           `json:"total_indexing_buffer_in_bytes,omitempty"`
	Transport                  *NodeInfoTransport `json:"transport,omitempty"`

	TransportAddress string `json:"transport_address"`
	TransportVersion int64  `json:"transport_version"`

	Version string `json:"version"`
}

func (s *NodeInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeInfo() *NodeInfo { _ = "STUB: not implemented"; return nil }
