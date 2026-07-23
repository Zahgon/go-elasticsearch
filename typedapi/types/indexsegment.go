package types

type IndexSegment struct {
	Shards map[string][]ShardsSegment `json:"shards"`
}

func (s *IndexSegment) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndexSegment() *IndexSegment { _ = "STUB: not implemented"; return nil }
