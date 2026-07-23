package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _doubleRangeProperty struct {
	v *types.DoubleRangeProperty
}

func NewDoubleRangeProperty() *_doubleRangeProperty { _ = "STUB: not implemented"; return nil }

func (s *_doubleRangeProperty) Boost(boost types.Float64) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) Coerce(coerce bool) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) CopyTo(fields ...string) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) DocValues(docvalues bool) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) Fields(fields map[string]types.Property) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) AddField(key string, value types.PropertyVariant) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) IgnoreAbove(ignoreabove int) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) Index(index bool) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) Meta(meta map[string]string) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) AddMeta(key string, value string) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) Properties(properties map[string]types.Property) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) AddProperty(key string, value types.PropertyVariant) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) Store(store bool) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_doubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleRangeProperty) DoubleRangePropertyCaster() *types.DoubleRangeProperty {
	_ = "STUB: not implemented"
	return nil
}
