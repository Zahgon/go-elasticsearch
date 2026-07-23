package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _floatRangeProperty struct {
	v *types.FloatRangeProperty
}

func NewFloatRangeProperty() *_floatRangeProperty { _ = "STUB: not implemented"; return nil }

func (s *_floatRangeProperty) Boost(boost types.Float64) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) Coerce(coerce bool) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) CopyTo(fields ...string) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) DocValues(docvalues bool) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) Fields(fields map[string]types.Property) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) AddField(key string, value types.PropertyVariant) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) IgnoreAbove(ignoreabove int) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) Index(index bool) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) Meta(meta map[string]string) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) AddMeta(key string, value string) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) Properties(properties map[string]types.Property) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) AddProperty(key string, value types.PropertyVariant) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) Store(store bool) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_floatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatRangeProperty) FloatRangePropertyCaster() *types.FloatRangeProperty {
	_ = "STUB: not implemented"
	return nil
}
