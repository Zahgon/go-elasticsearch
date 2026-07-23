package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type HistogramProperty struct {
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	TimeSeriesMetric    *timeseriesmetrictype.TimeSeriesMetricType       `json:"time_series_metric,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *HistogramProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s HistogramProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewHistogramProperty() *HistogramProperty { _ = "STUB: not implemented"; return nil }

type HistogramPropertyVariant interface {
	HistogramPropertyCaster() *HistogramProperty
}

func (s *HistogramProperty) HistogramPropertyCaster() *HistogramProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *HistogramProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
