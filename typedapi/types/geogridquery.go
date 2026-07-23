package types

type GeoGridQuery struct {
	Boost      *float32 `json:"boost,omitempty"`
	Geohash    *string  `json:"geohash,omitempty"`
	Geohex     *string  `json:"geohex,omitempty"`
	Geotile    *string  `json:"geotile,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`
}

func (s *GeoGridQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGeoGridQuery() *GeoGridQuery { _ = "STUB: not implemented"; return nil }

type GeoGridQueryVariant interface {
	GeoGridQueryCaster() *GeoGridQuery
}

func (s *GeoGridQuery) GeoGridQueryCaster() *GeoGridQuery { _ = "STUB: not implemented"; return nil }
