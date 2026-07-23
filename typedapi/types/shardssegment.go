package types

type ShardsSegment struct {
	NumCommittedSegments int                 `json:"num_committed_segments"`
	NumSearchSegments    int                 `json:"num_search_segments"`
	Routing              ShardSegmentRouting `json:"routing"`
	Segments             map[string]Segment  `json:"segments"`
}

func (s *ShardsSegment) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShardsSegment() *ShardsSegment { _ = "STUB: not implemented"; return nil }
