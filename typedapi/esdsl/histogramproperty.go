package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _histogramProperty struct {
	v *types.HistogramProperty
}

func NewHistogramProperty() *_histogramProperty { _ = "STUB: not implemented"; return nil }

func (s *_histogramProperty) IgnoreMalformed(ignoremalformed bool) *_histogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_histogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_histogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramProperty) Fields(fields map[string]types.Property) *_histogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramProperty) AddField(key string, value types.PropertyVariant) *_histogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramProperty) IgnoreAbove(ignoreabove int) *_histogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramProperty) Meta(meta map[string]string) *_histogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramProperty) AddMeta(key string, value string) *_histogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramProperty) Properties(properties map[string]types.Property) *_histogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramProperty) AddProperty(key string, value types.PropertyVariant) *_histogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_histogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramProperty) HistogramPropertyCaster() *types.HistogramProperty {
	_ = "STUB: not implemented"
	return nil
}
