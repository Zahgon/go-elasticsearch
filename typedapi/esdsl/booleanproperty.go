package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _booleanProperty struct {
	v *types.BooleanProperty
}

func NewBooleanProperty() *_booleanProperty { _ = "STUB: not implemented"; return nil }

func (s *_booleanProperty) Boost(boost types.Float64) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) Fielddata(fielddata types.NumericFielddataVariant) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) IgnoreMalformed(ignoremalformed bool) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) Index(index bool) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) NullValue(nullvalue bool) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) Script(script types.ScriptVariant) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) TimeSeriesDimension(timeseriesdimension bool) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) CopyTo(fields ...string) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) DocValues(docvalues bool) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) Fields(fields map[string]types.Property) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) AddField(key string, value types.PropertyVariant) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) IgnoreAbove(ignoreabove int) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) Meta(meta map[string]string) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) AddMeta(key string, value string) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) Properties(properties map[string]types.Property) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) AddProperty(key string, value types.PropertyVariant) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) Store(store bool) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_booleanProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_booleanProperty) BooleanPropertyCaster() *types.BooleanProperty {
	_ = "STUB: not implemented"
	return nil
}
