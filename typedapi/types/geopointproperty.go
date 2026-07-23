package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geopointmetrictype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type GeoPointProperty struct {
	CopyTo          []string                       `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	IgnoreZValue    *bool                          `json:"ignore_z_value,omitempty"`
	Index           *bool                          `json:"index,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	NullValue           GeoLocation                                      `json:"null_value,omitempty"`
	OnScriptError       *onscripterror.OnScriptError                     `json:"on_script_error,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Script              *Script                                          `json:"script,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	TimeSeriesMetric    *geopointmetrictype.GeoPointMetricType           `json:"time_series_metric,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *GeoPointProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GeoPointProperty) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGeoPointProperty() *GeoPointProperty { _ = "STUB: not implemented"; return nil }

type GeoPointPropertyVariant interface {
	GeoPointPropertyCaster() *GeoPointProperty
}

func (s *GeoPointProperty) GeoPointPropertyCaster() *GeoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *GeoPointProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
