package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoexecution"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geovalidationmethod"
)

type _geoBoundingBoxQuery struct {
	v *types.GeoBoundingBoxQuery
}

func NewGeoBoundingBoxQuery() *_geoBoundingBoxQuery { _ = "STUB: not implemented"; return nil }

func (s *_geoBoundingBoxQuery) GeoBoundingBoxQuery(geoboundingboxquery map[string]types.GeoBounds) *_geoBoundingBoxQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoBoundingBoxQuery) AddGeoBoundingBoxQuery(key string, value types.GeoBoundsVariant) *_geoBoundingBoxQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoBoundingBoxQuery) IgnoreUnmapped(ignoreunmapped bool) *_geoBoundingBoxQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoBoundingBoxQuery) Type(type_ geoexecution.GeoExecution) *_geoBoundingBoxQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoBoundingBoxQuery) ValidationMethod(validationmethod geovalidationmethod.GeoValidationMethod) *_geoBoundingBoxQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoBoundingBoxQuery) Boost(boost float32) *_geoBoundingBoxQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoBoundingBoxQuery) QueryName_(queryname_ string) *_geoBoundingBoxQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoBoundingBoxQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_geoBoundingBoxQuery) GeoBoundingBoxQueryCaster() *types.GeoBoundingBoxQuery {
	_ = "STUB: not implemented"
	return nil
}
