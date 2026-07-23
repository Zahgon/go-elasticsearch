package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _floatNumberProperty struct {
	v *types.FloatNumberProperty
}

func NewFloatNumberProperty() *_floatNumberProperty { _ = "STUB: not implemented"; return nil }

func (s *_floatNumberProperty) NullValue(nullvalue float32) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) Boost(boost types.Float64) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) Coerce(coerce bool) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) CopyTo(fields ...string) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) DocValues(docvalues bool) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) Fields(fields map[string]types.Property) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) AddField(key string, value types.PropertyVariant) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) IgnoreAbove(ignoreabove int) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) IgnoreMalformed(ignoremalformed bool) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) Index(index bool) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) Meta(meta map[string]string) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) AddMeta(key string, value string) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) Properties(properties map[string]types.Property) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) AddProperty(key string, value types.PropertyVariant) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) Script(script types.ScriptVariant) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) Store(store bool) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_floatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_floatNumberProperty) FloatNumberPropertyCaster() *types.FloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}
