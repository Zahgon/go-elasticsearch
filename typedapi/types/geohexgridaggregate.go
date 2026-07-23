package types

type GeoHexGridAggregate struct {
	Buckets BucketsGeoHexGridBucket `json:"buckets"`
	Meta    Metadata                `json:"meta,omitempty"`
}

func (s *GeoHexGridAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoHexGridAggregate() *GeoHexGridAggregate { _ = "STUB: not implemented"; return nil }
