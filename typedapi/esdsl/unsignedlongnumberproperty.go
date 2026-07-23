package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _unsignedLongNumberProperty struct {
	v *types.UnsignedLongNumberProperty
}

func NewUnsignedLongNumberProperty() *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) NullValue(nullvalue uint64) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) Boost(boost types.Float64) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) Coerce(coerce bool) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) CopyTo(fields ...string) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) DocValues(docvalues bool) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) Fields(fields map[string]types.Property) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) AddField(key string, value types.PropertyVariant) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) IgnoreAbove(ignoreabove int) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) IgnoreMalformed(ignoremalformed bool) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) Index(index bool) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) Meta(meta map[string]string) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) AddMeta(key string, value string) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) Properties(properties map[string]types.Property) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) AddProperty(key string, value types.PropertyVariant) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) Script(script types.ScriptVariant) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) Store(store bool) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_unsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_unsignedLongNumberProperty) UnsignedLongNumberPropertyCaster() *types.UnsignedLongNumberProperty {
	_ = "STUB: not implemented"
	return nil
}
