package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/densevectorelementtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/densevectorsimilarity"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _denseVectorProperty struct {
	v *types.DenseVectorProperty
}

func NewDenseVectorProperty() *_denseVectorProperty { _ = "STUB: not implemented"; return nil }

func (s *_denseVectorProperty) Dims(dims int) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) ElementType(elementtype densevectorelementtype.DenseVectorElementType) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) Index(index bool) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) IndexOptions(indexoptions types.DenseVectorIndexOptionsVariant) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) Similarity(similarity densevectorsimilarity.DenseVectorSimilarity) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) Fields(fields map[string]types.Property) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) AddField(key string, value types.PropertyVariant) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) IgnoreAbove(ignoreabove int) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) Meta(meta map[string]string) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) AddMeta(key string, value string) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) Properties(properties map[string]types.Property) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) AddProperty(key string, value types.PropertyVariant) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_denseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_denseVectorProperty) DenseVectorPropertyCaster() *types.DenseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}
