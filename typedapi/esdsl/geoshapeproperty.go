package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoorientation"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geostrategy"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _geoShapeProperty struct {
	v *types.GeoShapeProperty
}

func NewGeoShapeProperty() *_geoShapeProperty { _ = "STUB: not implemented"; return nil }

func (s *_geoShapeProperty) Coerce(coerce bool) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) IgnoreMalformed(ignoremalformed bool) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) IgnoreZValue(ignorezvalue bool) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) Index(index bool) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) Orientation(orientation geoorientation.GeoOrientation) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) Strategy(strategy geostrategy.GeoStrategy) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) CopyTo(fields ...string) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) DocValues(docvalues bool) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) Fields(fields map[string]types.Property) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) AddField(key string, value types.PropertyVariant) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) IgnoreAbove(ignoreabove int) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) Meta(meta map[string]string) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) AddMeta(key string, value string) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) Properties(properties map[string]types.Property) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) AddProperty(key string, value types.PropertyVariant) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) Store(store bool) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_geoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoShapeProperty) GeoShapePropertyCaster() *types.GeoShapeProperty {
	_ = "STUB: not implemented"
	return nil
}
