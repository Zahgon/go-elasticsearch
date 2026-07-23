package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shardstoreallocation"
)

type ShardStore struct {
	Allocation shardstoreallocation.ShardStoreAllocation `json:"allocation"`

	AllocationId *string                   `json:"allocation_id,omitempty"`
	ShardStore   map[string]ShardStoreNode `json:"-"`

	StoreException *ShardStoreException `json:"store_exception,omitempty"`
}

func (s *ShardStore) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ShardStore) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewShardStore() *ShardStore { _ = "STUB: not implemented"; return nil }
