package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _fieldAliasProperty struct {
	v *types.FieldAliasProperty
}

func NewFieldAliasProperty() *_fieldAliasProperty { _ = "STUB: not implemented"; return nil }

func (s *_fieldAliasProperty) Path(field string) *_fieldAliasProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAliasProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_fieldAliasProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAliasProperty) Fields(fields map[string]types.Property) *_fieldAliasProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAliasProperty) AddField(key string, value types.PropertyVariant) *_fieldAliasProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAliasProperty) IgnoreAbove(ignoreabove int) *_fieldAliasProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAliasProperty) Meta(meta map[string]string) *_fieldAliasProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAliasProperty) AddMeta(key string, value string) *_fieldAliasProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAliasProperty) Properties(properties map[string]types.Property) *_fieldAliasProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAliasProperty) AddProperty(key string, value types.PropertyVariant) *_fieldAliasProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAliasProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_fieldAliasProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAliasProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAliasProperty) FieldAliasPropertyCaster() *types.FieldAliasProperty {
	_ = "STUB: not implemented"
	return nil
}
