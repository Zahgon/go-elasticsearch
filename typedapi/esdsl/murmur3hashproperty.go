package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _murmur3HashProperty struct {
	v *types.Murmur3HashProperty
}

func NewMurmur3HashProperty() *_murmur3HashProperty { _ = "STUB: not implemented"; return nil }

func (s *_murmur3HashProperty) CopyTo(fields ...string) *_murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_murmur3HashProperty) DocValues(docvalues bool) *_murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_murmur3HashProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_murmur3HashProperty) Fields(fields map[string]types.Property) *_murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_murmur3HashProperty) AddField(key string, value types.PropertyVariant) *_murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_murmur3HashProperty) IgnoreAbove(ignoreabove int) *_murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_murmur3HashProperty) Meta(meta map[string]string) *_murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_murmur3HashProperty) AddMeta(key string, value string) *_murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_murmur3HashProperty) Properties(properties map[string]types.Property) *_murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_murmur3HashProperty) AddProperty(key string, value types.PropertyVariant) *_murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_murmur3HashProperty) Store(store bool) *_murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_murmur3HashProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_murmur3HashProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_murmur3HashProperty) Murmur3HashPropertyCaster() *types.Murmur3HashProperty {
	_ = "STUB: not implemented"
	return nil
}
