package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoShapeQuery struct {
	v *types.GeoShapeQuery
}

func NewGeoShapeQuery() *_geoShapeQuery { _ = "STUB: not implemented"; return nil }

func (s *_geoShapeQuery) GeoShapeQuery(geoshapequery map[string]types.GeoShapeFieldQuery) *_geoShapeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeQuery) AddGeoShapeQuery(key string, value types.GeoShapeFieldQueryVariant) *_geoShapeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeQuery) IgnoreUnmapped(ignoreunmapped bool) *_geoShapeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeQuery) Boost(boost float32) *_geoShapeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeQuery) QueryName_(queryname_ string) *_geoShapeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_geoShapeQuery) GeoShapeQueryCaster() *types.GeoShapeQuery {
	_ = "STUB: not implemented"
	return nil
}
