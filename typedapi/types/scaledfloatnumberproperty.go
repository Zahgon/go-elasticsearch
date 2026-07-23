package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type ScaledFloatNumberProperty struct {
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
	NullValue           *Float64                                         `json:"null_value,omitempty"`
	OnScriptError       *onscripterror.OnScriptError                     `json:"on_script_error,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	ScalingFactor       *Float64                                         `json:"scaling_factor,omitempty"`
	Script              *Script                                          `json:"script,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`

	TimeSeriesDimension *bool `json:"time_series_dimension,omitempty"`

	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type,omitempty"`
}

func (s *ScaledFloatNumberProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ScaledFloatNumberProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewScaledFloatNumberProperty() *ScaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

type ScaledFloatNumberPropertyVariant interface {
	ScaledFloatNumberPropertyCaster() *ScaledFloatNumberProperty
}

func (s *ScaledFloatNumberProperty) ScaledFloatNumberPropertyCaster() *ScaledFloatNumberProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *ScaledFloatNumberProperty) PropertyCaster() *Property {
	_ = "STUB: not implemented"
	return nil
}
