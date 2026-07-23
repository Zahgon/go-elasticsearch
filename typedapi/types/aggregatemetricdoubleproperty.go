package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type AggregateMetricDoubleProperty struct {
	DefaultMetric   string                         `json:"default_metric"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Metrics             []string                                         `json:"metrics"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	TimeSeriesMetric    *timeseriesmetrictype.TimeSeriesMetricType       `json:"time_series_metric,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *AggregateMetricDoubleProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s AggregateMetricDoubleProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewAggregateMetricDoubleProperty() *AggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

type AggregateMetricDoublePropertyVariant interface {
	AggregateMetricDoublePropertyCaster() *AggregateMetricDoubleProperty
}

func (s *AggregateMetricDoubleProperty) AggregateMetricDoublePropertyCaster() *AggregateMetricDoubleProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *AggregateMetricDoubleProperty) PropertyCaster() *Property {
	_ = "STUB: not implemented"
	return nil
}
