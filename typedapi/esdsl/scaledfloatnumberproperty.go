package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _scaledFloatNumberProperty struct {
	v *types.ScaledFloatNumberProperty
}

func NewScaledFloatNumberProperty() *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) NullValue(nullvalue types.Float64) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) ScalingFactor(scalingfactor types.Float64) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) Boost(boost types.Float64) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) Coerce(coerce bool) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) CopyTo(fields ...string) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) DocValues(docvalues bool) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) Fields(fields map[string]types.Property) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) AddField(key string, value types.PropertyVariant) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) IgnoreAbove(ignoreabove int) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) IgnoreMalformed(ignoremalformed bool) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) Index(index bool) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) Meta(meta map[string]string) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) AddMeta(key string, value string) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) Properties(properties map[string]types.Property) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) AddProperty(key string, value types.PropertyVariant) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) Script(script types.ScriptVariant) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) Store(store bool) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_scaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scaledFloatNumberProperty) ScaledFloatNumberPropertyCaster() *types.ScaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}
