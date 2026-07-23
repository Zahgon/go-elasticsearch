package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geovalidationmethod"
)

type _geoDistanceQuery struct {
	v *types.GeoDistanceQuery
}

func NewGeoDistanceQuery() *_geoDistanceQuery { _ = "STUB: not implemented"; return nil }

func (s *_geoDistanceQuery) Distance(distance string) *_geoDistanceQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceQuery) DistanceType(distancetype geodistancetype.GeoDistanceType) *_geoDistanceQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceQuery) GeoDistanceQuery(geodistancequery map[string]types.GeoLocation) *_geoDistanceQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceQuery) AddGeoDistanceQuery(key string, value types.GeoLocationVariant) *_geoDistanceQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceQuery) IgnoreUnmapped(ignoreunmapped bool) *_geoDistanceQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceQuery) ValidationMethod(validationmethod geovalidationmethod.GeoValidationMethod) *_geoDistanceQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceQuery) Boost(boost float32) *_geoDistanceQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceQuery) QueryName_(queryname_ string) *_geoDistanceQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDistanceQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_geoDistanceQuery) GeoDistanceQueryCaster() *types.GeoDistanceQuery {
	_ = "STUB: not implemented"
	return nil
}
