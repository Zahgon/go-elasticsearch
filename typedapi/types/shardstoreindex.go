package types

type ShardStoreIndex struct {
	Aliases []string `json:"aliases,omitempty"`
	Filter  *Query   `json:"filter,omitempty"`
}

func NewShardStoreIndex() *ShardStoreIndex { _ = "STUB: not implemented"; return nil }
