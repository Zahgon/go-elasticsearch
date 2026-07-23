package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _shortNumberProperty struct {
	v *types.ShortNumberProperty
}

func NewShortNumberProperty() *_shortNumberProperty { _ = "STUB: not implemented"; return nil }

func (s *_shortNumberProperty) NullValue(nullvalue int) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) Boost(boost types.Float64) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) Coerce(coerce bool) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) CopyTo(fields ...string) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) DocValues(docvalues bool) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) Fields(fields map[string]types.Property) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) AddField(key string, value types.PropertyVariant) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) IgnoreAbove(ignoreabove int) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) IgnoreMalformed(ignoremalformed bool) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) Index(index bool) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) Meta(meta map[string]string) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) AddMeta(key string, value string) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) Properties(properties map[string]types.Property) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) AddProperty(key string, value types.PropertyVariant) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) Script(script types.ScriptVariant) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) Store(store bool) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_shortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shortNumberProperty) ShortNumberPropertyCaster() *types.ShortNumberProperty {
	_ = "STUB: not implemented"
	return nil
}
