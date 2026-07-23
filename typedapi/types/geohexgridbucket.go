package types

type GeoHexGridBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Key          string               `json:"key"`
}

func (s *GeoHexGridBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GeoHexGridBucket) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGeoHexGridBucket() *GeoHexGridBucket { _ = "STUB: not implemented"; return nil }
