package types

type GeoLinePoint struct {
	Field string `json:"field"`
}

func (s *GeoLinePoint) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGeoLinePoint() *GeoLinePoint { _ = "STUB: not implemented"; return nil }

type GeoLinePointVariant interface {
	GeoLinePointCaster() *GeoLinePoint
}

func (s *GeoLinePoint) GeoLinePointCaster() *GeoLinePoint { _ = "STUB: not implemented"; return nil }
