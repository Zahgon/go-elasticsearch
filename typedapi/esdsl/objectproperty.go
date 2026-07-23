package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/subobjects"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _objectProperty struct {
	v *types.ObjectProperty
}

func NewObjectProperty() *_objectProperty { _ = "STUB: not implemented"; return nil }

func (s *_objectProperty) Enabled(enabled bool) *_objectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_objectProperty) Subobjects(subobjects subobjects.Subobjects) *_objectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_objectProperty) CopyTo(fields ...string) *_objectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_objectProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_objectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_objectProperty) Fields(fields map[string]types.Property) *_objectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_objectProperty) AddField(key string, value types.PropertyVariant) *_objectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_objectProperty) IgnoreAbove(ignoreabove int) *_objectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_objectProperty) Meta(meta map[string]string) *_objectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_objectProperty) AddMeta(key string, value string) *_objectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_objectProperty) Properties(properties map[string]types.Property) *_objectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_objectProperty) AddProperty(key string, value types.PropertyVariant) *_objectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_objectProperty) Store(store bool) *_objectProperty { _ = "STUB: not implemented"; return nil }

func (s *_objectProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_objectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_objectProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_objectProperty) ObjectPropertyCaster() *types.ObjectProperty {
	_ = "STUB: not implemented"
	return nil
}
