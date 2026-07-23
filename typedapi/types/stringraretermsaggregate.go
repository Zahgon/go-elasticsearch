package types

type StringRareTermsAggregate struct {
	Buckets BucketsStringRareTermsBucket `json:"buckets"`
	Meta    Metadata                     `json:"meta,omitempty"`
}

func (s *StringRareTermsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewStringRareTermsAggregate() *StringRareTermsAggregate { _ = "STUB: not implemented"; return nil }
