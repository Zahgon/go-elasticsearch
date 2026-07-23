package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoorientation"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _shapeProperty struct {
	v *types.ShapeProperty
}

func NewShapeProperty() *_shapeProperty { _ = "STUB: not implemented"; return nil }

func (s *_shapeProperty) Coerce(coerce bool) *_shapeProperty { _ = "STUB: not implemented"; return nil }

func (s *_shapeProperty) IgnoreMalformed(ignoremalformed bool) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) IgnoreZValue(ignorezvalue bool) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) Orientation(orientation geoorientation.GeoOrientation) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) CopyTo(fields ...string) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) DocValues(docvalues bool) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) Fields(fields map[string]types.Property) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) AddField(key string, value types.PropertyVariant) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) IgnoreAbove(ignoreabove int) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) Meta(meta map[string]string) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) AddMeta(key string, value string) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) Properties(properties map[string]types.Property) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) AddProperty(key string, value types.PropertyVariant) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) Store(store bool) *_shapeProperty { _ = "STUB: not implemented"; return nil }

func (s *_shapeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_shapeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeProperty) ShapePropertyCaster() *types.ShapeProperty {
	_ = "STUB: not implemented"
	return nil
}
