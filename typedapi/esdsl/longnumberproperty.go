package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _longNumberProperty struct {
	v *types.LongNumberProperty
}

func NewLongNumberProperty() *_longNumberProperty { _ = "STUB: not implemented"; return nil }

func (s *_longNumberProperty) NullValue(nullvalue int64) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) Boost(boost types.Float64) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) Coerce(coerce bool) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) CopyTo(fields ...string) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) DocValues(docvalues bool) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) Fields(fields map[string]types.Property) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) AddField(key string, value types.PropertyVariant) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) IgnoreAbove(ignoreabove int) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) IgnoreMalformed(ignoremalformed bool) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) Index(index bool) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) Meta(meta map[string]string) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) AddMeta(key string, value string) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) Properties(properties map[string]types.Property) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) AddProperty(key string, value types.PropertyVariant) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) Script(script types.ScriptVariant) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) Store(store bool) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_longNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberProperty) LongNumberPropertyCaster() *types.LongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}
