package types

type ShardStoreWrapper struct {
	Stores []ShardStore `json:"stores"`
}

func NewShardStoreWrapper() *ShardStoreWrapper { _ = "STUB: not implemented"; return nil }
