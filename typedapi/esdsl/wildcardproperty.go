package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _wildcardProperty struct {
	v *types.WildcardProperty
}

func NewWildcardProperty() *_wildcardProperty { _ = "STUB: not implemented"; return nil }

func (s *_wildcardProperty) NullValue(nullvalue string) *_wildcardProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) CopyTo(fields ...string) *_wildcardProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) DocValues(docvalues bool) *_wildcardProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_wildcardProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) Fields(fields map[string]types.Property) *_wildcardProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) AddField(key string, value types.PropertyVariant) *_wildcardProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) IgnoreAbove(ignoreabove int) *_wildcardProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) Meta(meta map[string]string) *_wildcardProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) AddMeta(key string, value string) *_wildcardProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) Properties(properties map[string]types.Property) *_wildcardProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) AddProperty(key string, value types.PropertyVariant) *_wildcardProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) Store(store bool) *_wildcardProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_wildcardProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardProperty) WildcardPropertyCaster() *types.WildcardProperty {
	_ = "STUB: not implemented"
	return nil
}
