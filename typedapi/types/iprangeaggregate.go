package types

type IpRangeAggregate struct {
	Buckets BucketsIpRangeBucket `json:"buckets"`
	Meta    Metadata             `json:"meta,omitempty"`
}

func (s *IpRangeAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIpRangeAggregate() *IpRangeAggregate { _ = "STUB: not implemented"; return nil }
