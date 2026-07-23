package types

type GeoTileGridBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Key          string               `json:"key"`
}

func (s *GeoTileGridBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GeoTileGridBucket) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGeoTileGridBucket() *GeoTileGridBucket { _ = "STUB: not implemented"; return nil }
