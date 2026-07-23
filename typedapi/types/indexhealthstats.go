package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/healthstatus"
)

type IndexHealthStats struct {
	ActivePrimaryShards     int                         `json:"active_primary_shards"`
	ActiveShards            int                         `json:"active_shards"`
	InitializingShards      int                         `json:"initializing_shards"`
	NumberOfReplicas        int                         `json:"number_of_replicas"`
	NumberOfShards          int                         `json:"number_of_shards"`
	RelocatingShards        int                         `json:"relocating_shards"`
	Shards                  map[string]ShardHealthStats `json:"shards,omitempty"`
	Status                  healthstatus.HealthStatus   `json:"status"`
	UnassignedPrimaryShards int                         `json:"unassigned_primary_shards"`
	UnassignedShards        int                         `json:"unassigned_shards"`
}

func (s *IndexHealthStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndexHealthStats() *IndexHealthStats { _ = "STUB: not implemented"; return nil }
