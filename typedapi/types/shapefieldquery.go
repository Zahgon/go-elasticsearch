package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoshaperelation"
)

type ShapeFieldQuery struct {
	IndexedShape *FieldLookup `json:"indexed_shape,omitempty"`

	Relation *geoshaperelation.GeoShapeRelation `json:"relation,omitempty"`

	Shape json.RawMessage `json:"shape,omitempty"`
}

func (s *ShapeFieldQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShapeFieldQuery() *ShapeFieldQuery { _ = "STUB: not implemented"; return nil }

type ShapeFieldQueryVariant interface {
	ShapeFieldQueryCaster() *ShapeFieldQuery
}

func (s *ShapeFieldQuery) ShapeFieldQueryCaster() *ShapeFieldQuery {
	_ = "STUB: not implemented"
	return nil
}
