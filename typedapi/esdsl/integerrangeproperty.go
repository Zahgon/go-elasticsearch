package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _integerRangeProperty struct {
	v *types.IntegerRangeProperty
}

func NewIntegerRangeProperty() *_integerRangeProperty { _ = "STUB: not implemented"; return nil }

func (s *_integerRangeProperty) Boost(boost types.Float64) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) Coerce(coerce bool) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) CopyTo(fields ...string) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) DocValues(docvalues bool) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) Fields(fields map[string]types.Property) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) AddField(key string, value types.PropertyVariant) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) IgnoreAbove(ignoreabove int) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) Index(index bool) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) Meta(meta map[string]string) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) AddMeta(key string, value string) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) Properties(properties map[string]types.Property) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) AddProperty(key string, value types.PropertyVariant) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) Store(store bool) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_integerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerRangeProperty) IntegerRangePropertyCaster() *types.IntegerRangeProperty {
	_ = "STUB: not implemented"
	return nil
}
