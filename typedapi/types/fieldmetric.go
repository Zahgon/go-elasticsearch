package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/metric"
)

type FieldMetric struct {
	Field string `json:"field"`

	Metrics []metric.Metric `json:"metrics"`
}

func (s *FieldMetric) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldMetric() *FieldMetric { _ = "STUB: not implemented"; return nil }

type FieldMetricVariant interface {
	FieldMetricCaster() *FieldMetric
}

func (s *FieldMetric) FieldMetricCaster() *FieldMetric { _ = "STUB: not implemented"; return nil }
