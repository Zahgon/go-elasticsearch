package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoshaperelation"
)

type _geoShapeFieldQuery struct {
	v *types.GeoShapeFieldQuery
}

func NewGeoShapeFieldQuery() *_geoShapeFieldQuery { _ = "STUB: not implemented"; return nil }

func (s *_geoShapeFieldQuery) IndexedShape(indexedshape types.FieldLookupVariant) *_geoShapeFieldQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeFieldQuery) Relation(relation geoshaperelation.GeoShapeRelation) *_geoShapeFieldQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeFieldQuery) Shape(geoshape json.RawMessage) *_geoShapeFieldQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeFieldQuery) GeoShapeFieldQueryCaster() *types.GeoShapeFieldQuery {
	_ = "STUB: not implemented"
	return nil
}
