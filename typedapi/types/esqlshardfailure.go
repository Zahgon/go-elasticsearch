package types

type EsqlShardFailure struct {
	Index  *string    `json:"index,omitempty"`
	Node   *string    `json:"node,omitempty"`
	Reason ErrorCause `json:"reason"`
	Shard  int        `json:"shard"`
}

func (s *EsqlShardFailure) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEsqlShardFailure() *EsqlShardFailure { _ = "STUB: not implemented"; return nil }
