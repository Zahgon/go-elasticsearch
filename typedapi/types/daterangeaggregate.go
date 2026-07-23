package types

type DateRangeAggregate struct {
	Buckets BucketsRangeBucket `json:"buckets"`
	Meta    Metadata           `json:"meta,omitempty"`
}

func (s *DateRangeAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDateRangeAggregate() *DateRangeAggregate { _ = "STUB: not implemented"; return nil }
