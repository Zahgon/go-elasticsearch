package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _ipProperty struct {
	v *types.IpProperty
}

func NewIpProperty() *_ipProperty { _ = "STUB: not implemented"; return nil }

func (s *_ipProperty) Boost(boost types.Float64) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) IgnoreMalformed(ignoremalformed bool) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) Index(index bool) *_ipProperty { _ = "STUB: not implemented"; return nil }

func (s *_ipProperty) NullValue(nullvalue string) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) Script(script types.ScriptVariant) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) TimeSeriesDimension(timeseriesdimension bool) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) CopyTo(fields ...string) *_ipProperty { _ = "STUB: not implemented"; return nil }

func (s *_ipProperty) DocValues(docvalues bool) *_ipProperty { _ = "STUB: not implemented"; return nil }

func (s *_ipProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) Fields(fields map[string]types.Property) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) AddField(key string, value types.PropertyVariant) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) IgnoreAbove(ignoreabove int) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) Meta(meta map[string]string) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) AddMeta(key string, value string) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) Properties(properties map[string]types.Property) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) AddProperty(key string, value types.PropertyVariant) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) Store(store bool) *_ipProperty { _ = "STUB: not implemented"; return nil }

func (s *_ipProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_ipProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipProperty) IpPropertyCaster() *types.IpProperty { _ = "STUB: not implemented"; return nil }
