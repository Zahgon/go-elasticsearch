package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _sparseVectorProperty struct {
	v *types.SparseVectorProperty
}

func NewSparseVectorProperty() *_sparseVectorProperty { _ = "STUB: not implemented"; return nil }

func (s *_sparseVectorProperty) IndexOptions(indexoptions types.SparseVectorIndexOptionsVariant) *_sparseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorProperty) Store(store bool) *_sparseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_sparseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorProperty) Fields(fields map[string]types.Property) *_sparseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorProperty) AddField(key string, value types.PropertyVariant) *_sparseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorProperty) IgnoreAbove(ignoreabove int) *_sparseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorProperty) Meta(meta map[string]string) *_sparseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorProperty) AddMeta(key string, value string) *_sparseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorProperty) Properties(properties map[string]types.Property) *_sparseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorProperty) AddProperty(key string, value types.PropertyVariant) *_sparseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_sparseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorProperty) SparseVectorPropertyCaster() *types.SparseVectorProperty {
	_ = "STUB: not implemented"
	return nil
}
