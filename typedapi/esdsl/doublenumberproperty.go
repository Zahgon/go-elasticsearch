package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _doubleNumberProperty struct {
	v *types.DoubleNumberProperty
}

func NewDoubleNumberProperty() *_doubleNumberProperty { _ = "STUB: not implemented"; return nil }

func (s *_doubleNumberProperty) NullValue(nullvalue types.Float64) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) Boost(boost types.Float64) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) Coerce(coerce bool) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) CopyTo(fields ...string) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) DocValues(docvalues bool) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) Fields(fields map[string]types.Property) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) AddField(key string, value types.PropertyVariant) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) IgnoreAbove(ignoreabove int) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) IgnoreMalformed(ignoremalformed bool) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) Index(index bool) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) Meta(meta map[string]string) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) AddMeta(key string, value string) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) Properties(properties map[string]types.Property) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) AddProperty(key string, value types.PropertyVariant) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) Script(script types.ScriptVariant) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) Store(store bool) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_doubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_doubleNumberProperty) DoubleNumberPropertyCaster() *types.DoubleNumberProperty {
	_ = "STUB: not implemented"
	return nil
}
