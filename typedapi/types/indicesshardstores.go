package types

type IndicesShardStores struct {
	Shards map[string]ShardStoreWrapper `json:"shards"`
}

func NewIndicesShardStores() *IndicesShardStores { _ = "STUB: not implemented"; return nil }
