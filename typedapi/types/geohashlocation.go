package types

type GeoHashLocation struct {
	Geohash string `json:"geohash"`
}

func (s *GeoHashLocation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGeoHashLocation() *GeoHashLocation { _ = "STUB: not implemented"; return nil }

type GeoHashLocationVariant interface {
	GeoHashLocationCaster() *GeoHashLocation
}

func (s *GeoHashLocation) GeoHashLocationCaster() *GeoHashLocation {
	_ = "STUB: not implemented"
	return nil
}

func (s *GeoHashLocation) GeoLocationCaster() *GeoLocation { _ = "STUB: not implemented"; return nil }
