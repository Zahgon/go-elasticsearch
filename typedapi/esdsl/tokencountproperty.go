package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _tokenCountProperty struct {
	v *types.TokenCountProperty
}

func NewTokenCountProperty() *_tokenCountProperty { _ = "STUB: not implemented"; return nil }

func (s *_tokenCountProperty) Analyzer(analyzer string) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) Boost(boost types.Float64) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) EnablePositionIncrements(enablepositionincrements bool) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) Index(index bool) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) NullValue(nullvalue types.Float64) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) CopyTo(fields ...string) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) DocValues(docvalues bool) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) Fields(fields map[string]types.Property) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) AddField(key string, value types.PropertyVariant) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) IgnoreAbove(ignoreabove int) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) Meta(meta map[string]string) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) AddMeta(key string, value string) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) Properties(properties map[string]types.Property) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) AddProperty(key string, value types.PropertyVariant) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) Store(store bool) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_tokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenCountProperty) TokenCountPropertyCaster() *types.TokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}
