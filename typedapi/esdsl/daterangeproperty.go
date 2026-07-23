package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _dateRangeProperty struct {
	v *types.DateRangeProperty
}

func NewDateRangeProperty() *_dateRangeProperty { _ = "STUB: not implemented"; return nil }

func (s *_dateRangeProperty) Format(format string) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) Boost(boost types.Float64) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) Coerce(coerce bool) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) CopyTo(fields ...string) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) DocValues(docvalues bool) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) Fields(fields map[string]types.Property) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) AddField(key string, value types.PropertyVariant) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) IgnoreAbove(ignoreabove int) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) Index(index bool) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) Meta(meta map[string]string) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) AddMeta(key string, value string) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) Properties(properties map[string]types.Property) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) AddProperty(key string, value types.PropertyVariant) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) Store(store bool) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_dateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeProperty) DateRangePropertyCaster() *types.DateRangeProperty {
	_ = "STUB: not implemented"
	return nil
}
