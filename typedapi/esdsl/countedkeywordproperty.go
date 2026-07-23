package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _countedKeywordProperty struct {
	v *types.CountedKeywordProperty
}

func NewCountedKeywordProperty() *_countedKeywordProperty { _ = "STUB: not implemented"; return nil }

func (s *_countedKeywordProperty) Index(index bool) *_countedKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_countedKeywordProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_countedKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_countedKeywordProperty) Fields(fields map[string]types.Property) *_countedKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_countedKeywordProperty) AddField(key string, value types.PropertyVariant) *_countedKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_countedKeywordProperty) IgnoreAbove(ignoreabove int) *_countedKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_countedKeywordProperty) Meta(meta map[string]string) *_countedKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_countedKeywordProperty) AddMeta(key string, value string) *_countedKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_countedKeywordProperty) Properties(properties map[string]types.Property) *_countedKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_countedKeywordProperty) AddProperty(key string, value types.PropertyVariant) *_countedKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_countedKeywordProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_countedKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_countedKeywordProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_countedKeywordProperty) CountedKeywordPropertyCaster() *types.CountedKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}
