package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geovalidationmethod"
)

type GeoPolygonQuery struct {
	Boost            *float32                                 `json:"boost,omitempty"`
	GeoPolygonQuery  map[string]GeoPolygonPoints              `json:"-"`
	IgnoreUnmapped   *bool                                    `json:"ignore_unmapped,omitempty"`
	QueryName_       *string                                  `json:"_name,omitempty"`
	ValidationMethod *geovalidationmethod.GeoValidationMethod `json:"validation_method,omitempty"`
}

func (s *GeoPolygonQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GeoPolygonQuery) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGeoPolygonQuery() *GeoPolygonQuery { _ = "STUB: not implemented"; return nil }

type GeoPolygonQueryVariant interface {
	GeoPolygonQueryCaster() *GeoPolygonQuery
}

func (s *GeoPolygonQuery) GeoPolygonQueryCaster() *GeoPolygonQuery {
	_ = "STUB: not implemented"
	return nil
}
