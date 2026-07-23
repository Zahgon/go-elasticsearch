package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _integerNumberProperty struct {
	v *types.IntegerNumberProperty
}

func NewIntegerNumberProperty() *_integerNumberProperty { _ = "STUB: not implemented"; return nil }

func (s *_integerNumberProperty) NullValue(nullvalue int) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) Boost(boost types.Float64) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) Coerce(coerce bool) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) CopyTo(fields ...string) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) DocValues(docvalues bool) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) Fields(fields map[string]types.Property) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) AddField(key string, value types.PropertyVariant) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) IgnoreAbove(ignoreabove int) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) IgnoreMalformed(ignoremalformed bool) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) Index(index bool) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) Meta(meta map[string]string) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) AddMeta(key string, value string) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) Properties(properties map[string]types.Property) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) AddProperty(key string, value types.PropertyVariant) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) Script(script types.ScriptVariant) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) Store(store bool) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_integerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_integerNumberProperty) IntegerNumberPropertyCaster() *types.IntegerNumberProperty {
	_ = "STUB: not implemented"
	return nil
}
