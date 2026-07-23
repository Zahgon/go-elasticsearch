package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type IpProperty struct {
	Boost           *Float64                       `json:"boost,omitempty"`
	CopyTo          []string                       `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	Index           *bool                          `json:"index,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	NullValue           *string                                          `json:"null_value,omitempty"`
	OnScriptError       *onscripterror.OnScriptError                     `json:"on_script_error,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Script              *Script                                          `json:"script,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`

	TimeSeriesDimension *bool  `json:"time_series_dimension,omitempty"`
	Type                string `json:"type,omitempty"`
}

func (s *IpProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s IpProperty) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewIpProperty() *IpProperty { _ = "STUB: not implemented"; return nil }

type IpPropertyVariant interface {
	IpPropertyCaster() *IpProperty
}

func (s *IpProperty) IpPropertyCaster() *IpProperty { _ = "STUB: not implemented"; return nil }

func (s *IpProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
