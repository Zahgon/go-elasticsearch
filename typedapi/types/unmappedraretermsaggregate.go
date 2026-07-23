package types

type UnmappedRareTermsAggregate struct {
	Buckets BucketsVoid `json:"buckets"`
	Meta    Metadata    `json:"meta,omitempty"`
}

func (s *UnmappedRareTermsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUnmappedRareTermsAggregate() *UnmappedRareTermsAggregate {
	_ = "STUB: not implemented"
	return nil
}
