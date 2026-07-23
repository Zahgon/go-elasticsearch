package types

type ShardSequenceNumber struct {
	GlobalCheckpoint int64 `json:"global_checkpoint"`
	LocalCheckpoint  int64 `json:"local_checkpoint"`
	MaxSeqNo         int64 `json:"max_seq_no"`
}

func (s *ShardSequenceNumber) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewShardSequenceNumber() *ShardSequenceNumber { _ = "STUB: not implemented"; return nil }
