package types

type ShardStoreException struct {
	Reason string `json:"reason"`
	Type   string `json:"type"`
}

func (s *ShardStoreException) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewShardStoreException() *ShardStoreException { _ = "STUB: not implemented"; return nil }
