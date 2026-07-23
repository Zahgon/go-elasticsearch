package types

type GeoDistanceFeatureQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Field string `json:"field"`

	Origin GeoLocation `json:"origin"`

	Pivot      string  `json:"pivot"`
	QueryName_ *string `json:"_name,omitempty"`
}

func (s *GeoDistanceFeatureQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoDistanceFeatureQuery() *GeoDistanceFeatureQuery { _ = "STUB: not implemented"; return nil }

type GeoDistanceFeatureQueryVariant interface {
	GeoDistanceFeatureQueryCaster() *GeoDistanceFeatureQuery
}

func (s *GeoDistanceFeatureQuery) GeoDistanceFeatureQueryCaster() *GeoDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *GeoDistanceFeatureQuery) DistanceFeatureQueryCaster() *DistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}
