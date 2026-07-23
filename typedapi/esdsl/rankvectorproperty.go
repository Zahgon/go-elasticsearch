package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rankvectorelementtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _rankVectorProperty struct {
	v *types.RankVectorProperty
}

func NewRankVectorProperty() *_rankVectorProperty { _ = "STUB: not implemented"; return nil }

func (s *_rankVectorProperty) Dims(dims int) *_rankVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankVectorProperty) ElementType(elementtype rankvectorelementtype.RankVectorElementType) *_rankVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankVectorProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_rankVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankVectorProperty) Fields(fields map[string]types.Property) *_rankVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankVectorProperty) AddField(key string, value types.PropertyVariant) *_rankVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankVectorProperty) IgnoreAbove(ignoreabove int) *_rankVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankVectorProperty) Meta(meta map[string]string) *_rankVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankVectorProperty) AddMeta(key string, value string) *_rankVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankVectorProperty) Properties(properties map[string]types.Property) *_rankVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankVectorProperty) AddProperty(key string, value types.PropertyVariant) *_rankVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankVectorProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_rankVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankVectorProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankVectorProperty) RankVectorPropertyCaster() *types.RankVectorProperty {
	_ = "STUB: not implemented"
	return nil
}
