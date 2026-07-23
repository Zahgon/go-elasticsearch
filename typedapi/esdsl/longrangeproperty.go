package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _longRangeProperty struct {
	v *types.LongRangeProperty
}

func NewLongRangeProperty() *_longRangeProperty { _ = "STUB: not implemented"; return nil }

func (s *_longRangeProperty) Boost(boost types.Float64) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) Coerce(coerce bool) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) CopyTo(fields ...string) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) DocValues(docvalues bool) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) Fields(fields map[string]types.Property) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) AddField(key string, value types.PropertyVariant) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) IgnoreAbove(ignoreabove int) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) Index(index bool) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) Meta(meta map[string]string) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) AddMeta(key string, value string) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) Properties(properties map[string]types.Property) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) AddProperty(key string, value types.PropertyVariant) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) Store(store bool) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_longRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longRangeProperty) LongRangePropertyCaster() *types.LongRangeProperty {
	_ = "STUB: not implemented"
	return nil
}
