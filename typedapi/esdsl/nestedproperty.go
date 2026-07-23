package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _nestedProperty struct {
	v *types.NestedProperty
}

func NewNestedProperty() *_nestedProperty { _ = "STUB: not implemented"; return nil }

func (s *_nestedProperty) Enabled(enabled bool) *_nestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) IncludeInParent(includeinparent bool) *_nestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) IncludeInRoot(includeinroot bool) *_nestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) CopyTo(fields ...string) *_nestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_nestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) Fields(fields map[string]types.Property) *_nestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) AddField(key string, value types.PropertyVariant) *_nestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) IgnoreAbove(ignoreabove int) *_nestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) Meta(meta map[string]string) *_nestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) AddMeta(key string, value string) *_nestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) Properties(properties map[string]types.Property) *_nestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) AddProperty(key string, value types.PropertyVariant) *_nestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) Store(store bool) *_nestedProperty { _ = "STUB: not implemented"; return nil }

func (s *_nestedProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_nestedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedProperty) NestedPropertyCaster() *types.NestedProperty {
	_ = "STUB: not implemented"
	return nil
}
