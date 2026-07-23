package types

type GeoHashGridBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Key          string               `json:"key"`
}

func (s *GeoHashGridBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GeoHashGridBucket) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGeoHashGridBucket() *GeoHashGridBucket { _ = "STUB: not implemented"; return nil }
