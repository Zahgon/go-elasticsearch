package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoorientation"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geostrategy"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type GeoShapeProperty struct {
	Coerce          *bool                          `json:"coerce,omitempty"`
	CopyTo          []string                       `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	IgnoreZValue    *bool                          `json:"ignore_z_value,omitempty"`
	Index           *bool                          `json:"index,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Orientation         *geoorientation.GeoOrientation                   `json:"orientation,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	Strategy            *geostrategy.GeoStrategy                         `json:"strategy,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *GeoShapeProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GeoShapeProperty) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGeoShapeProperty() *GeoShapeProperty { _ = "STUB: not implemented"; return nil }

type GeoShapePropertyVariant interface {
	GeoShapePropertyCaster() *GeoShapeProperty
}

func (s *GeoShapeProperty) GeoShapePropertyCaster() *GeoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *GeoShapeProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
