package types

type GeoPolygonPoints struct {
	Points []GeoLocation `json:"points"`
}

func NewGeoPolygonPoints() *GeoPolygonPoints { _ = "STUB: not implemented"; return nil }

type GeoPolygonPointsVariant interface {
	GeoPolygonPointsCaster() *GeoPolygonPoints
}

func (s *GeoPolygonPoints) GeoPolygonPointsCaster() *GeoPolygonPoints {
	_ = "STUB: not implemented"
	return nil
}
