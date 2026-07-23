package types

type GeoTileGridAggregate struct {
	Buckets BucketsGeoTileGridBucket `json:"buckets"`
	Meta    Metadata                 `json:"meta,omitempty"`
}

func (s *GeoTileGridAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoTileGridAggregate() *GeoTileGridAggregate { _ = "STUB: not implemented"; return nil }
