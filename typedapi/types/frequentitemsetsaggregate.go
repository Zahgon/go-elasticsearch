package types

type FrequentItemSetsAggregate struct {
	Buckets BucketsFrequentItemSetsBucket `json:"buckets"`
	Meta    Metadata                      `json:"meta,omitempty"`
}

func (s *FrequentItemSetsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFrequentItemSetsAggregate() *FrequentItemSetsAggregate {
	_ = "STUB: not implemented"
	return nil
}
