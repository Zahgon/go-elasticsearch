package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _percolatorProperty struct {
	v *types.PercolatorProperty
}

func NewPercolatorProperty() *_percolatorProperty { _ = "STUB: not implemented"; return nil }

func (s *_percolatorProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_percolatorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolatorProperty) Fields(fields map[string]types.Property) *_percolatorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolatorProperty) AddField(key string, value types.PropertyVariant) *_percolatorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolatorProperty) IgnoreAbove(ignoreabove int) *_percolatorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolatorProperty) Meta(meta map[string]string) *_percolatorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolatorProperty) AddMeta(key string, value string) *_percolatorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolatorProperty) Properties(properties map[string]types.Property) *_percolatorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolatorProperty) AddProperty(key string, value types.PropertyVariant) *_percolatorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolatorProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_percolatorProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolatorProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolatorProperty) PercolatorPropertyCaster() *types.PercolatorProperty {
	_ = "STUB: not implemented"
	return nil
}
