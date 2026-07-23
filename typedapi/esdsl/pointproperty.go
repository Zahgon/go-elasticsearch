package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _pointProperty struct {
	v *types.PointProperty
}

func NewPointProperty() *_pointProperty { _ = "STUB: not implemented"; return nil }

func (s *_pointProperty) IgnoreMalformed(ignoremalformed bool) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) IgnoreZValue(ignorezvalue bool) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) NullValue(nullvalue string) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) CopyTo(fields ...string) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) DocValues(docvalues bool) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) Fields(fields map[string]types.Property) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) AddField(key string, value types.PropertyVariant) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) IgnoreAbove(ignoreabove int) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) Meta(meta map[string]string) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) AddMeta(key string, value string) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) Properties(properties map[string]types.Property) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) AddProperty(key string, value types.PropertyVariant) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) Store(store bool) *_pointProperty { _ = "STUB: not implemented"; return nil }

func (s *_pointProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_pointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointProperty) PointPropertyCaster() *types.PointProperty {
	_ = "STUB: not implemented"
	return nil
}
