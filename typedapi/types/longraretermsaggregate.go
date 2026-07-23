package types

type LongRareTermsAggregate struct {
	Buckets BucketsLongRareTermsBucket `json:"buckets"`
	Meta    Metadata                   `json:"meta,omitempty"`
}

func (s *LongRareTermsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewLongRareTermsAggregate() *LongRareTermsAggregate { _ = "STUB: not implemented"; return nil }
