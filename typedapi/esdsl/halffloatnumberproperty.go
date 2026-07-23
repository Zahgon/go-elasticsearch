package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _halfFloatNumberProperty struct {
	v *types.HalfFloatNumberProperty
}

func NewHalfFloatNumberProperty() *_halfFloatNumberProperty { _ = "STUB: not implemented"; return nil }

func (s *_halfFloatNumberProperty) NullValue(nullvalue float32) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) Boost(boost types.Float64) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) Coerce(coerce bool) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) CopyTo(fields ...string) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) DocValues(docvalues bool) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) Fields(fields map[string]types.Property) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) AddField(key string, value types.PropertyVariant) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) IgnoreAbove(ignoreabove int) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) IgnoreMalformed(ignoremalformed bool) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) Index(index bool) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) Meta(meta map[string]string) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) AddMeta(key string, value string) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) Properties(properties map[string]types.Property) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) AddProperty(key string, value types.PropertyVariant) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) Script(script types.ScriptVariant) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) Store(store bool) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_halfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_halfFloatNumberProperty) HalfFloatNumberPropertyCaster() *types.HalfFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}
