package types

type ShardFailure struct {
	Index   *string    `json:"index,omitempty"`
	Node    *string    `json:"node,omitempty"`
	Primary *bool      `json:"primary,omitempty"`
	Reason  ErrorCause `json:"reason"`
	Shard   *int       `json:"shard,omitempty"`
	Status  *string    `json:"status,omitempty"`
}

func (s *ShardFailure) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShardFailure() *ShardFailure { _ = "STUB: not implemented"; return nil }
