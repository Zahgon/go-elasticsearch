package types

type SignificantLongTermsAggregate struct {
	BgCount  *int64                            `json:"bg_count,omitempty"`
	Buckets  BucketsSignificantLongTermsBucket `json:"buckets"`
	DocCount *int64                            `json:"doc_count,omitempty"`
	Meta     Metadata                          `json:"meta,omitempty"`
}

func (s *SignificantLongTermsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSignificantLongTermsAggregate() *SignificantLongTermsAggregate {
	_ = "STUB: not implemented"
	return nil
}
