package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _aggregateMetricDoubleProperty struct {
	v *types.AggregateMetricDoubleProperty
}

func NewAggregateMetricDoubleProperty(defaultmetric string) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) DefaultMetric(defaultmetric string) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) IgnoreMalformed(ignoremalformed bool) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) Metrics(metrics ...string) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) Fields(fields map[string]types.Property) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) AddField(key string, value types.PropertyVariant) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) IgnoreAbove(ignoreabove int) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) Meta(meta map[string]string) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) AddMeta(key string, value string) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) Properties(properties map[string]types.Property) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) AddProperty(key string, value types.PropertyVariant) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_aggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregateMetricDoubleProperty) AggregateMetricDoublePropertyCaster() *types.AggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}
