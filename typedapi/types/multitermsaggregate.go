package types

type MultiTermsAggregate struct {
	Buckets                 BucketsMultiTermsBucket `json:"buckets"`
	DocCountErrorUpperBound *int64                  `json:"doc_count_error_upper_bound,omitempty"`
	Meta                    Metadata                `json:"meta,omitempty"`
	SumOtherDocCount        *int64                  `json:"sum_other_doc_count,omitempty"`
}

func (s *MultiTermsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMultiTermsAggregate() *MultiTermsAggregate { _ = "STUB: not implemented"; return nil }
