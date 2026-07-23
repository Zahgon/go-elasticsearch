package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _versionProperty struct {
	v *types.VersionProperty
}

func NewVersionProperty() *_versionProperty { _ = "STUB: not implemented"; return nil }

func (s *_versionProperty) CopyTo(fields ...string) *_versionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_versionProperty) DocValues(docvalues bool) *_versionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_versionProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_versionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_versionProperty) Fields(fields map[string]types.Property) *_versionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_versionProperty) AddField(key string, value types.PropertyVariant) *_versionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_versionProperty) IgnoreAbove(ignoreabove int) *_versionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_versionProperty) Meta(meta map[string]string) *_versionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_versionProperty) AddMeta(key string, value string) *_versionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_versionProperty) Properties(properties map[string]types.Property) *_versionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_versionProperty) AddProperty(key string, value types.PropertyVariant) *_versionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_versionProperty) Store(store bool) *_versionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_versionProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_versionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_versionProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_versionProperty) VersionPropertyCaster() *types.VersionProperty {
	_ = "STUB: not implemented"
	return nil
}
