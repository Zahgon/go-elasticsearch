package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geovalidationmethod"
)

type _geoPolygonQuery struct {
	v *types.GeoPolygonQuery
}

func NewGeoPolygonQuery() *_geoPolygonQuery { _ = "STUB: not implemented"; return nil }

func (s *_geoPolygonQuery) GeoPolygonQuery(geopolygonquery map[string]types.GeoPolygonPoints) *_geoPolygonQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPolygonQuery) AddGeoPolygonQuery(key string, value types.GeoPolygonPointsVariant) *_geoPolygonQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPolygonQuery) IgnoreUnmapped(ignoreunmapped bool) *_geoPolygonQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPolygonQuery) ValidationMethod(validationmethod geovalidationmethod.GeoValidationMethod) *_geoPolygonQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPolygonQuery) Boost(boost float32) *_geoPolygonQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPolygonQuery) QueryName_(queryname_ string) *_geoPolygonQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPolygonQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_geoPolygonQuery) GeoPolygonQueryCaster() *types.GeoPolygonQuery {
	_ = "STUB: not implemented"
	return nil
}
