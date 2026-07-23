package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoshaperelation"
)

type GeoShapeFieldQuery struct {
	IndexedShape *FieldLookup `json:"indexed_shape,omitempty"`

	Relation *geoshaperelation.GeoShapeRelation `json:"relation,omitempty"`
	Shape    json.RawMessage                    `json:"shape,omitempty"`
}

func (s *GeoShapeFieldQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoShapeFieldQuery() *GeoShapeFieldQuery { _ = "STUB: not implemented"; return nil }

type GeoShapeFieldQueryVariant interface {
	GeoShapeFieldQueryCaster() *GeoShapeFieldQuery
}

func (s *GeoShapeFieldQuery) GeoShapeFieldQueryCaster() *GeoShapeFieldQuery {
	_ = "STUB: not implemented"
	return nil
}
