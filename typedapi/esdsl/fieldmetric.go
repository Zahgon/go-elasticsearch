package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/metric"
)

type _fieldMetric struct {
	v *types.FieldMetric
}

func NewFieldMetric() *_fieldMetric { _ = "STUB: not implemented"; return nil }

func (s *_fieldMetric) Field(field string) *_fieldMetric { _ = "STUB: not implemented"; return nil }

func (s *_fieldMetric) Metrics(metrics ...metric.Metric) *_fieldMetric {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldMetric) FieldMetricCaster() *types.FieldMetric {
	_ = "STUB: not implemented"
	return nil
}
