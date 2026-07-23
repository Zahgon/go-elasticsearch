package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type ExponentialHistogramProperty struct {
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	TimeSeriesMetric    *timeseriesmetrictype.TimeSeriesMetricType       `json:"time_series_metric,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *ExponentialHistogramProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ExponentialHistogramProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewExponentialHistogramProperty() *ExponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

type ExponentialHistogramPropertyVariant interface {
	ExponentialHistogramPropertyCaster() *ExponentialHistogramProperty
}

func (s *ExponentialHistogramProperty) ExponentialHistogramPropertyCaster() *ExponentialHistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *ExponentialHistogramProperty) PropertyCaster() *Property {
	_ = "STUB: not implemented"
	return nil
}
