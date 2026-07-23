package types

type GeoDistanceAggregate struct {
	Buckets BucketsRangeBucket `json:"buckets"`
	Meta    Metadata           `json:"meta,omitempty"`
}

func (s *GeoDistanceAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoDistanceAggregate() *GeoDistanceAggregate { _ = "STUB: not implemented"; return nil }
