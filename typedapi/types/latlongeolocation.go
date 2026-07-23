package types

type LatLonGeoLocation struct {
	Lat Float64 `json:"lat"`

	Lon Float64 `json:"lon"`
}

func (s *LatLonGeoLocation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewLatLonGeoLocation() *LatLonGeoLocation { _ = "STUB: not implemented"; return nil }

type LatLonGeoLocationVariant interface {
	LatLonGeoLocationCaster() *LatLonGeoLocation
}

func (s *LatLonGeoLocation) LatLonGeoLocationCaster() *LatLonGeoLocation {
	_ = "STUB: not implemented"
	return nil
}

func (s *LatLonGeoLocation) GeoLocationCaster() *GeoLocation { _ = "STUB: not implemented"; return nil }
