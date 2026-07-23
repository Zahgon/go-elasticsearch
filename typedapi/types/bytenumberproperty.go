package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type ByteNumberProperty struct {
	Boost           *Float64                       `json:"boost,omitempty"`
	Coerce          *bool                          `json:"coerce,omitempty"`
	CopyTo          []string                       `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	Index           *bool                          `json:"index,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	NullValue           *byte                                            `json:"null_value,omitempty"`
	OnScriptError       *onscripterror.OnScriptError                     `json:"on_script_error,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Script              *Script                                          `json:"script,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`

	TimeSeriesDimension *bool `json:"time_series_dimension,omitempty"`

	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type,omitempty"`
}

func (s *ByteNumberProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ByteNumberProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewByteNumberProperty() *ByteNumberProperty { _ = "STUB: not implemented"; return nil }

type ByteNumberPropertyVariant interface {
	ByteNumberPropertyCaster() *ByteNumberProperty
}

func (s *ByteNumberProperty) ByteNumberPropertyCaster() *ByteNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *ByteNumberProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
