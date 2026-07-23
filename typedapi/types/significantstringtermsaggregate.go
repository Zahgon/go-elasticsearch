package types

type SignificantStringTermsAggregate struct {
	BgCount  *int64                              `json:"bg_count,omitempty"`
	Buckets  BucketsSignificantStringTermsBucket `json:"buckets"`
	DocCount *int64                              `json:"doc_count,omitempty"`
	Meta     Metadata                            `json:"meta,omitempty"`
}

func (s *SignificantStringTermsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSignificantStringTermsAggregate() *SignificantStringTermsAggregate {
	_ = "STUB: not implemented"
	return nil
}
