package types

type RangeAggregate struct {
	Buckets BucketsRangeBucket `json:"buckets"`
	Meta    Metadata           `json:"meta,omitempty"`
}

func (s *RangeAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRangeAggregate() *RangeAggregate { _ = "STUB: not implemented"; return nil }
