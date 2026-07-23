package types

type IndexDetails struct {
	MaxSegmentsPerShard int64    `json:"max_segments_per_shard"`
	ShardCount          int      `json:"shard_count"`
	Size                ByteSize `json:"size,omitempty"`
	SizeInBytes         int64    `json:"size_in_bytes"`
}

func (s *IndexDetails) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndexDetails() *IndexDetails { _ = "STUB: not implemented"; return nil }
