package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _exponentialHistogramProperty struct {
	v *types.ExponentialHistogramProperty
}

func NewExponentialHistogramProperty() *_exponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exponentialHistogramProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_exponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exponentialHistogramProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_exponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exponentialHistogramProperty) Fields(fields map[string]types.Property) *_exponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exponentialHistogramProperty) AddField(key string, value types.PropertyVariant) *_exponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exponentialHistogramProperty) IgnoreAbove(ignoreabove int) *_exponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exponentialHistogramProperty) Meta(meta map[string]string) *_exponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exponentialHistogramProperty) AddMeta(key string, value string) *_exponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exponentialHistogramProperty) Properties(properties map[string]types.Property) *_exponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exponentialHistogramProperty) AddProperty(key string, value types.PropertyVariant) *_exponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exponentialHistogramProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_exponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exponentialHistogramProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_exponentialHistogramProperty) ExponentialHistogramPropertyCaster() *types.ExponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}
