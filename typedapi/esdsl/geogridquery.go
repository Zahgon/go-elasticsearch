package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoGridQuery struct {
	v *types.GeoGridQuery
}

func NewGeoGridQuery() *_geoGridQuery { _ = "STUB: not implemented"; return nil }

func (s *_geoGridQuery) Geohash(geohash string) *_geoGridQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridQuery) Geohex(geohexcell string) *_geoGridQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridQuery) Geotile(geotile string) *_geoGridQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridQuery) Boost(boost float32) *_geoGridQuery { _ = "STUB: not implemented"; return nil }

func (s *_geoGridQuery) QueryName_(queryname_ string) *_geoGridQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridQuery) GeoGridQueryCaster() *types.GeoGridQuery {
	_ = "STUB: not implemented"
	return nil
}
