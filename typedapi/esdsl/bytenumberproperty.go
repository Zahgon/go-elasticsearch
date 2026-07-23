package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _byteNumberProperty struct {
	v *types.ByteNumberProperty
}

func NewByteNumberProperty() *_byteNumberProperty { _ = "STUB: not implemented"; return nil }

func (s *_byteNumberProperty) NullValue(nullvalue byte) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) Boost(boost types.Float64) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) Coerce(coerce bool) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) CopyTo(fields ...string) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) DocValues(docvalues bool) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) Fields(fields map[string]types.Property) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) AddField(key string, value types.PropertyVariant) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) IgnoreAbove(ignoreabove int) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) IgnoreMalformed(ignoremalformed bool) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) Index(index bool) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) Meta(meta map[string]string) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) AddMeta(key string, value string) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) Properties(properties map[string]types.Property) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) AddProperty(key string, value types.PropertyVariant) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) Script(script types.ScriptVariant) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) Store(store bool) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_byteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_byteNumberProperty) ByteNumberPropertyCaster() *types.ByteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}
