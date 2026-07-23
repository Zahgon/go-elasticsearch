package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/healthstatus"
)

type ShardHealthStats struct {
	ActiveShards            int                       `json:"active_shards"`
	InitializingShards      int                       `json:"initializing_shards"`
	PrimaryActive           bool                      `json:"primary_active"`
	RelocatingShards        int                       `json:"relocating_shards"`
	Status                  healthstatus.HealthStatus `json:"status"`
	UnassignedPrimaryShards int                       `json:"unassigned_primary_shards"`
	UnassignedShards        int                       `json:"unassigned_shards"`
}

func (s *ShardHealthStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShardHealthStats() *ShardHealthStats { _ = "STUB: not implemented"; return nil }
