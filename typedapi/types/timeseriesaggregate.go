package types

type TimeSeriesAggregate struct {
	Buckets BucketsTimeSeriesBucket `json:"buckets"`
	Meta    Metadata                `json:"meta,omitempty"`
}

func (s *TimeSeriesAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTimeSeriesAggregate() *TimeSeriesAggregate { _ = "STUB: not implemented"; return nil }
