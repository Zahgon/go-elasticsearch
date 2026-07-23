package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/noderole"
)

type Stats struct {
	AdaptiveSelection map[string]AdaptiveSelection `json:"adaptive_selection,omitempty"`

	Allocations *Allocations `json:"allocations,omitempty"`

	Attributes map[string]string `json:"attributes,omitempty"`

	Breakers map[string]Breaker `json:"breakers,omitempty"`

	Discovery *Discovery `json:"discovery,omitempty"`

	Fs *FileSystem `json:"fs,omitempty"`

	Host *string `json:"host,omitempty"`

	Http *Http `json:"http,omitempty"`

	IndexingPressure *NodesIndexingPressure `json:"indexing_pressure,omitempty"`

	Indices *IndicesShardStats `json:"indices,omitempty"`

	Ingest *NodesIngest `json:"ingest,omitempty"`

	Ip []string `json:"ip,omitempty"`

	Jvm *Jvm `json:"jvm,omitempty"`

	Name *string `json:"name,omitempty"`

	Os *OperatingSystem `json:"os,omitempty"`

	Process *Process `json:"process,omitempty"`

	Roles []noderole.NodeRole `json:"roles,omitempty"`

	Script      *Scripting               `json:"script,omitempty"`
	ScriptCache map[string][]ScriptCache `json:"script_cache,omitempty"`

	ThreadPool map[string]ThreadCount `json:"thread_pool,omitempty"`
	Timestamp  *int64                 `json:"timestamp,omitempty"`

	Transport *Transport `json:"transport,omitempty"`

	TransportAddress *string `json:"transport_address,omitempty"`
}

func (s *Stats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewStats() *Stats { _ = "STUB: not implemented"; return nil }
