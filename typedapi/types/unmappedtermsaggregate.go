package types

type UnmappedTermsAggregate struct {
	Buckets                 BucketsVoid `json:"buckets"`
	DocCountErrorUpperBound *int64      `json:"doc_count_error_upper_bound,omitempty"`
	Meta                    Metadata    `json:"meta,omitempty"`
	SumOtherDocCount        *int64      `json:"sum_other_doc_count,omitempty"`
}

func (s *UnmappedTermsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUnmappedTermsAggregate() *UnmappedTermsAggregate { _ = "STUB: not implemented"; return nil }
