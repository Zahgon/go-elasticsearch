package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/healthstatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexmetadatastate"
)

type IndicesStats struct {
	Health    *healthstatus.HealthStatus             `json:"health,omitempty"`
	Primaries *IndexStats                            `json:"primaries,omitempty"`
	Shards    map[string][]IndicesShardStats         `json:"shards,omitempty"`
	Status    *indexmetadatastate.IndexMetadataState `json:"status,omitempty"`
	Total     *IndexStats                            `json:"total,omitempty"`
	Uuid      *string                                `json:"uuid,omitempty"`
}

func (s *IndicesStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndicesStats() *IndicesStats { _ = "STUB: not implemented"; return nil }
