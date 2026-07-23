package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _dateProperty struct {
	v *types.DateProperty
}

func NewDateProperty() *_dateProperty { _ = "STUB: not implemented"; return nil }

func (s *_dateProperty) Boost(boost types.Float64) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) Fielddata(fielddata types.NumericFielddataVariant) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) Format(format string) *_dateProperty { _ = "STUB: not implemented"; return nil }

func (s *_dateProperty) IgnoreMalformed(ignoremalformed bool) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) Index(index bool) *_dateProperty { _ = "STUB: not implemented"; return nil }

func (s *_dateProperty) Locale(locale string) *_dateProperty { _ = "STUB: not implemented"; return nil }

func (s *_dateProperty) NullValue(datetime types.DateTimeVariant) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) PrecisionStep(precisionstep int) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) Script(script types.ScriptVariant) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) CopyTo(fields ...string) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) DocValues(docvalues bool) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) Fields(fields map[string]types.Property) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) AddField(key string, value types.PropertyVariant) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) IgnoreAbove(ignoreabove int) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) Meta(meta map[string]string) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) AddMeta(key string, value string) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) Properties(properties map[string]types.Property) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) AddProperty(key string, value types.PropertyVariant) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) Store(store bool) *_dateProperty { _ = "STUB: not implemented"; return nil }

func (s *_dateProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_dateProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProperty) DatePropertyCaster() *types.DateProperty {
	_ = "STUB: not implemented"
	return nil
}
