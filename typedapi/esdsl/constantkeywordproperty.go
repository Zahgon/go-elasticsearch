package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _constantKeywordProperty struct {
	v *types.ConstantKeywordProperty
}

func NewConstantKeywordProperty() *_constantKeywordProperty { _ = "STUB: not implemented"; return nil }

func (s *_constantKeywordProperty) Value(value json.RawMessage) *_constantKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantKeywordProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_constantKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantKeywordProperty) Fields(fields map[string]types.Property) *_constantKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantKeywordProperty) AddField(key string, value types.PropertyVariant) *_constantKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantKeywordProperty) IgnoreAbove(ignoreabove int) *_constantKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantKeywordProperty) Meta(meta map[string]string) *_constantKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantKeywordProperty) AddMeta(key string, value string) *_constantKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantKeywordProperty) Properties(properties map[string]types.Property) *_constantKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantKeywordProperty) AddProperty(key string, value types.PropertyVariant) *_constantKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantKeywordProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_constantKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantKeywordProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantKeywordProperty) ConstantKeywordPropertyCaster() *types.ConstantKeywordProperty {
	_ = "STUB: not implemented"
	return nil
}
