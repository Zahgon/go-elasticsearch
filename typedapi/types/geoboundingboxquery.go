package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoexecution"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geovalidationmethod"
)

type GeoBoundingBoxQuery struct {
	Boost               *float32             `json:"boost,omitempty"`
	GeoBoundingBoxQuery map[string]GeoBounds `json:"-"`

	IgnoreUnmapped *bool                      `json:"ignore_unmapped,omitempty"`
	QueryName_     *string                    `json:"_name,omitempty"`
	Type           *geoexecution.GeoExecution `json:"type,omitempty"`

	ValidationMethod *geovalidationmethod.GeoValidationMethod `json:"validation_method,omitempty"`
}

func (s *GeoBoundingBoxQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s GeoBoundingBoxQuery) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGeoBoundingBoxQuery() *GeoBoundingBoxQuery { _ = "STUB: not implemented"; return nil }

type GeoBoundingBoxQueryVariant interface {
	GeoBoundingBoxQueryCaster() *GeoBoundingBoxQuery
}

func (s *GeoBoundingBoxQuery) GeoBoundingBoxQueryCaster() *GeoBoundingBoxQuery {
	_ = "STUB: not implemented"
	return nil
}
