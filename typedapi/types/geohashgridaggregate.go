package types

type GeoHashGridAggregate struct {
	Buckets BucketsGeoHashGridBucket `json:"buckets"`
	Meta    Metadata                 `json:"meta,omitempty"`
}

func (s *GeoHashGridAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoHashGridAggregate() *GeoHashGridAggregate { _ = "STUB: not implemented"; return nil }
