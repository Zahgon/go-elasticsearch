package types

type ShardLease struct {
	Id             string `json:"id"`
	RetainingSeqNo int64  `json:"retaining_seq_no"`
	Source         string `json:"source"`
	Timestamp      int64  `json:"timestamp"`
}

func (s *ShardLease) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShardLease() *ShardLease { _ = "STUB: not implemented"; return nil }
