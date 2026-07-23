package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _dateNanosProperty struct {
	v *types.DateNanosProperty
}

func NewDateNanosProperty() *_dateNanosProperty { _ = "STUB: not implemented"; return nil }

func (s *_dateNanosProperty) Boost(boost types.Float64) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) Format(format string) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) IgnoreMalformed(ignoremalformed bool) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) Index(index bool) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) NullValue(datetime types.DateTimeVariant) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) PrecisionStep(precisionstep int) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) Script(script types.ScriptVariant) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) CopyTo(fields ...string) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) DocValues(docvalues bool) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) Fields(fields map[string]types.Property) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) AddField(key string, value types.PropertyVariant) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) IgnoreAbove(ignoreabove int) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) Meta(meta map[string]string) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) AddMeta(key string, value string) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) Properties(properties map[string]types.Property) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) AddProperty(key string, value types.PropertyVariant) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) Store(store bool) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_dateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateNanosProperty) DateNanosPropertyCaster() *types.DateNanosProperty {
	_ = "STUB: not implemented"
	return nil
}
