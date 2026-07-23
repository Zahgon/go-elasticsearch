package types

type ShardSegmentRouting struct {
	Node string `json:"node"`

	Primary bool `json:"primary"`

	State string `json:"state"`
}

func (s *ShardSegmentRouting) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewShardSegmentRouting() *ShardSegmentRouting { _ = "STUB: not implemented"; return nil }
