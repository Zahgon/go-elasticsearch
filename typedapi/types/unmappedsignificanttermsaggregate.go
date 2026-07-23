package types

type UnmappedSignificantTermsAggregate struct {
	BgCount  *int64      `json:"bg_count,omitempty"`
	Buckets  BucketsVoid `json:"buckets"`
	DocCount *int64      `json:"doc_count,omitempty"`
	Meta     Metadata    `json:"meta,omitempty"`
}

func (s *UnmappedSignificantTermsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUnmappedSignificantTermsAggregate() *UnmappedSignificantTermsAggregate {
	_ = "STUB: not implemented"
	return nil
}
