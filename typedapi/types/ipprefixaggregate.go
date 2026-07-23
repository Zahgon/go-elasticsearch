package types

type IpPrefixAggregate struct {
	Buckets BucketsIpPrefixBucket `json:"buckets"`
	Meta    Metadata              `json:"meta,omitempty"`
}

func (s *IpPrefixAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIpPrefixAggregate() *IpPrefixAggregate { _ = "STUB: not implemented"; return nil }
