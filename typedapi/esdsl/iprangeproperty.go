package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _ipRangeProperty struct {
	v *types.IpRangeProperty
}

func NewIpRangeProperty() *_ipRangeProperty { _ = "STUB: not implemented"; return nil }

func (s *_ipRangeProperty) Boost(boost types.Float64) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) Coerce(coerce bool) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) CopyTo(fields ...string) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) DocValues(docvalues bool) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) Fields(fields map[string]types.Property) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) AddField(key string, value types.PropertyVariant) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) IgnoreAbove(ignoreabove int) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) Index(index bool) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) Meta(meta map[string]string) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) AddMeta(key string, value string) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) Properties(properties map[string]types.Property) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) AddProperty(key string, value types.PropertyVariant) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) Store(store bool) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_ipRangeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipRangeProperty) IpRangePropertyCaster() *types.IpRangeProperty {
	_ = "STUB: not implemented"
	return nil
}
