package types

type GeoIpNodeDatabaseName struct {
	Name string `json:"name"`
}

func (s *GeoIpNodeDatabaseName) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoIpNodeDatabaseName() *GeoIpNodeDatabaseName { _ = "STUB: not implemented"; return nil }
