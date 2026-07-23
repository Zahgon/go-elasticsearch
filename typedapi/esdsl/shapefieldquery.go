package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoshaperelation"
)

type _shapeFieldQuery struct {
	v *types.ShapeFieldQuery
}

func NewShapeFieldQuery() *_shapeFieldQuery { _ = "STUB: not implemented"; return nil }

func (s *_shapeFieldQuery) IndexedShape(indexedshape types.FieldLookupVariant) *_shapeFieldQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeFieldQuery) Relation(relation geoshaperelation.GeoShapeRelation) *_shapeFieldQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeFieldQuery) Shape(geoshape json.RawMessage) *_shapeFieldQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeFieldQuery) ShapeFieldQueryCaster() *types.ShapeFieldQuery {
	_ = "STUB: not implemented"
	return nil
}
